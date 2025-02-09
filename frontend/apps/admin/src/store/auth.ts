import type { Recordable, UserInfo } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { DEFAULT_HOME_PATH, LOGIN_PATH } from '@vben/constants';
import { preferences } from '@vben/preferences';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';

import { notification } from 'ant-design-vue';
import { defineStore } from 'pinia';

import { $t } from '#/locales';
import { defAuthnService, defRouterService } from '#/rpc';

export const useAuthStore = defineStore('auth', () => {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const router = useRouter();

  const loginLoading = ref(false);

  /**
   * 异步处理登录操作
   * Asynchronously handle the login process
   * @param params 登录表单数据
   * @param onSuccess
   */
  async function authLogin(
    params: Recordable<any>,
    onSuccess?: () => Promise<void> | void,
  ) {
    // 异步处理用户登录操作并获取 accessToken
    let userInfo: null | UserInfo = null;
    try {
      loginLoading.value = true;

      const { access_token } = await defAuthnService.Login({
        username: params.username,
        password: params.password,
        grant_type: 'password',
      });

      // 如果成功获取到 accessToken
      if (access_token) {
        accessStore.setAccessToken(access_token);

        // 获取用户信息并存储到 accessStore 中
        const [fetchUserInfoResult, accessCodes] = await Promise.all([
          fetchUserInfo(),
          fetchAccessCodes(),
        ]);

        userInfo = fetchUserInfoResult;

        userStore.setUserInfo(userInfo);
        accessStore.setAccessCodes(accessCodes.codes);

        if (accessStore.loginExpired) {
          accessStore.setLoginExpired(false);
        } else {
          onSuccess
            ? await onSuccess?.()
            : await router.push(userInfo.homePath || DEFAULT_HOME_PATH);
        }

        if (userInfo?.realName) {
          notification.success({
            description: `${$t('authentication.loginSuccessDesc')}:${userInfo?.realName}`,
            duration: 3,
            message: $t('authentication.loginSuccess'),
          });
        }
      }
    } finally {
      loginLoading.value = false;
    }

    return {
      userInfo,
    };
  }

  /**
   * 用户登出
   * @param redirect
   */
  async function logout(redirect: boolean = true) {
    try {
      await defAuthnService.Logout({ id: 0 });
    } catch {
      // 不做任何处理
    }
    resetAllStores();
    accessStore.setLoginExpired(false);

    // 回登录页带上当前路由地址
    await router.replace({
      path: LOGIN_PATH,
      query: redirect
        ? {
            redirect: encodeURIComponent(router.currentRoute.value.fullPath),
          }
        : {},
    });
  }

  /**
   * 刷新访问令牌
   */
  async function refreshToken() {
    const accessStore = useAccessStore();

    const resp = await defAuthnService.RefreshToken({
      grant_type: 'password',
      refresh_token: accessStore.refreshToken ?? '',
    });
    const newToken = resp.access_token;

    accessStore.setAccessToken(newToken);

    return newToken;
  }

  /**
   * 重新认证
   */
  async function reauthenticate() {
    console.warn('Access token or refresh token is invalid or expired. ');
    const accessStore = useAccessStore();

    accessStore.setAccessToken(null);
    if (
      preferences.app.loginExpiredMode === 'modal' &&
      accessStore.isAccessChecked
    ) {
      accessStore.setLoginExpired(true);
    } else {
      await logout();
    }
  }

  /**
   * 拉取用户信息
   */
  async function fetchUserInfo() {
    return (await defAuthnService.GetMe({ id: 0 })) as UserInfo;
  }

  /**
   * 获取用户权限码
   */
  async function fetchAccessCodes() {
    return await defRouterService.ListPermissionCode({});
  }

  function $reset() {
    loginLoading.value = false;
  }

  return {
    $reset,
    authLogin,
    fetchUserInfo,
    loginLoading,
    logout,
    refreshToken,
    reauthenticate,
  };
});
