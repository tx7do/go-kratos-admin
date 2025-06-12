import type { AuthenticationService } from '#/generated/api/admin/service/v1/i_authentication.pb';
import type {
  ChangePasswordRequest,
  LoginRequest,
  LoginResponse,
} from '#/generated/api/authentication/service/v1/authentication.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';

import { requestClient } from '#/utils/request';

export type { AuthenticationService } from '#/generated/api/admin/service/v1/i_authentication.pb';
export type {
  LoginRequest,
  LoginResponse,
} from '#/generated/api/authentication/service/v1/authentication.pb';

/** 用户后台登录认证服务 */
export class AuthenticationServiceImpl implements AuthenticationService {
  async ChangePassword(request: ChangePasswordRequest): Promise<Empty> {
    return requestClient.post<Empty>('/change_password', request);
  }

  async Login(request: LoginRequest): Promise<LoginResponse> {
    return await requestClient.post<LoginResponse>('/login', request);
  }

  async Logout(_request: Empty): Promise<Empty> {
    return await requestClient.post<Empty>('/logout', {});
  }

  async RefreshToken(request: LoginRequest): Promise<LoginResponse> {
    return requestClient.post<LoginResponse>('/refresh_token', request);
  }
}

export const defAuthnService = new AuthenticationServiceImpl();
