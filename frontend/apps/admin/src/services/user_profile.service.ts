import type { UserProfileService } from '#/generated/api/admin/service/v1/i_user_profile.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type {
  UpdateUserRequest,
  User,
} from '#/generated/api/user/service/v1/user.pb';

import { requestClient } from '#/utils/request';

export type { UserProfileService } from '#/generated/api/admin/service/v1/i_user_profile.pb';
export type { User } from '#/generated/api/user/service/v1/user.pb';

/** 用户个人资料服务 */
export class UserProfileServiceImpl implements UserProfileService {
  async GetUser(_request: Empty): Promise<User> {
    return await requestClient.get<User>('/me');
  }

  async UpdateUser(request: UpdateUserRequest): Promise<Empty> {
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/me`, request);
  }
}

export const defUserProfileService = new UserProfileServiceImpl();
