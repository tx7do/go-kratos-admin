// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: admin/service/v1/i_admin_login_log.proto

/* eslint-disable */
import type { Timestamp } from "../../../google/protobuf/timestamp.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 后台登录日志 */
export interface AdminLoginLog {
  /** 后台登录日志ID */
  id?:
    | number
    | null
    | undefined;
  /** 登录IP地址 */
  loginIp?:
    | string
    | null
    | undefined;
  /** 登录MAC地址 */
  loginMac?:
    | string
    | null
    | undefined;
  /** 登录时间 */
  loginTime?:
    | Timestamp
    | null
    | undefined;
  /** 状态码 */
  statusCode?:
    | number
    | null
    | undefined;
  /** 登录是否成功 */
  success?:
    | boolean
    | null
    | undefined;
  /** 登录失败原因 */
  reason?:
    | string
    | null
    | undefined;
  /** 登录地理位置 */
  location?:
    | string
    | null
    | undefined;
  /** 浏览器的用户代理信息 */
  userAgent?:
    | string
    | null
    | undefined;
  /** 浏览器名称 */
  browserName?:
    | string
    | null
    | undefined;
  /** 浏览器版本 */
  browserVersion?:
    | string
    | null
    | undefined;
  /** 客户端ID */
  clientId?:
    | string
    | null
    | undefined;
  /** 客户端名称 */
  clientName?:
    | string
    | null
    | undefined;
  /** 操作系统名称 */
  osName?:
    | string
    | null
    | undefined;
  /** 操作系统版本 */
  osVersion?:
    | string
    | null
    | undefined;
  /** 操作者用户ID */
  userId?:
    | number
    | null
    | undefined;
  /** 操作者账号名 */
  username?:
    | string
    | null
    | undefined;
  /** 创建时间 */
  createTime?: Timestamp | null | undefined;
}

/** 查询后台登录日志列表 - 回应 */
export interface ListAdminLoginLogResponse {
  items: AdminLoginLog[];
  total: number;
}

/** 查询后台登录日志详情 - 请求 */
export interface GetAdminLoginLogRequest {
  id: number;
}

/** 创建后台登录日志 - 请求 */
export interface CreateAdminLoginLogRequest {
  data: AdminLoginLog | null;
}

/** 更新后台登录日志 - 请求 */
export interface UpdateAdminLoginLogRequest {
  data:
    | AdminLoginLog
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除后台登录日志 - 请求 */
export interface DeleteAdminLoginLogRequest {
  id: number;
}

/** 后台登录日志管理服务 */
export interface AdminLoginLogService {
  /** 查询后台登录日志列表 */
  List(request: PagingRequest): Promise<ListAdminLoginLogResponse>;
  /** 查询后台登录日志详情 */
  Get(request: GetAdminLoginLogRequest): Promise<AdminLoginLog>;
}
