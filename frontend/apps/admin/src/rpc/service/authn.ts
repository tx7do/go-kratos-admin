import type { AuthenticationService } from '#/rpc/api/admin/service/v1/i_authentication.pb';
import type {
  GetMeRequest,
  LoginRequest,
  LoginResponse,
  LogoutRequest,
} from '#/rpc/api/authentication/service/v1/authentication.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { User } from '#/rpc/api/user/service/v1/user.pb';

import { requestClient } from '#/rpc/request';

export type { AuthenticationService } from '#/rpc/api/admin/service/v1/i_authentication.pb';
export type {
  GetMeRequest,
  LoginRequest,
  LoginResponse,
  LogoutRequest,
} from '#/rpc/api/authentication/service/v1/authentication.pb';

/** 用户后台登录认证服务 */
export class AuthenticationServiceImpl implements AuthenticationService {
  async GetMe(_request: Empty): Promise<User> {
    return await requestClient.get<User>('/me');
  }

  async Login(request: LoginRequest): Promise<LoginResponse> {
    return await requestClient.post<LoginResponse>('/login', request);
  }

  async Logout(_request: Empty): Promise<Empty> {
    return await requestClient.post('/logout');
  }

  async RefreshToken(request: LoginRequest): Promise<LoginResponse> {
    return requestClient.post<LoginResponse>('/refresh_token', request);
  }
}

export const defAuthnService = new AuthenticationServiceImpl();
