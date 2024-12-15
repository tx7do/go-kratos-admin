interface BasicOption {
  label: string;
  value: string;
}

type SelectOption = BasicOption;

type TabOption = BasicOption;

interface BasicUserInfo {
  /**
   * 头像
   */
  avatar: string;
  /**
   * 用户id
   */
  id: number;
  /**
   * 用户昵称
   */
  nickName: string;
  /**
   * 用户真名
   */
  realName: string;
  /**
   * 用户角色
   */
  roles?: string[];
  /**
   * 用户名
   */
  userName: string;
}

type ClassType = Array<object | string> | object | string;

export type { BasicOption, BasicUserInfo, ClassType, SelectOption, TabOption };
