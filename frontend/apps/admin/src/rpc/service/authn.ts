import type {
  AuthenticationService,
  GetMeRequest,
  LoginRequest,
  LoginResponse,
  LogoutRequest,
  RefreshTokenRequest,
} from '#/rpc/api/admin/service/v1/i_authentication.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { User } from '#/rpc/api/user/service/v1/user.pb';

import { baseRequestClient, requestClient } from '#/rpc/request';

export type {
  AuthenticationService,
  GetMeRequest,
  LoginRequest,
  LoginResponse,
  LogoutRequest,
  RefreshTokenRequest,
} from '#/rpc/api/admin/service/v1/i_authentication.pb';

/** 用户后台登录认证服务 */
export class AuthenticationServiceImpl implements AuthenticationService {
  async GetMe(_request: GetMeRequest): Promise<User> {
    return await requestClient.get<User>('/me');
  }

  async Login(request: LoginRequest): Promise<LoginResponse> {
    return await requestClient.post<LoginResponse>('/login', request);
  }

  async Logout(_request: LogoutRequest): Promise<Empty> {
    return await baseRequestClient.post('/logout', {
      withCredentials: true,
    });
  }

  async RefreshToken(request: RefreshTokenRequest): Promise<LoginResponse> {
    return baseRequestClient.post<LoginResponse>('/refresh_token', request);
  }
}

export const defAuthnService = new AuthenticationServiceImpl();
