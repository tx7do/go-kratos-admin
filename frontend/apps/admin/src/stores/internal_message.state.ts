import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  createInternalMessageRecipientServiceClient,
  createInternalMessageServiceClient,
  type internal_messageservicev1_InternalMessage_Status as InternalMessage_Status,
  type internal_messageservicev1_InternalMessage_Type as InternalMessage_Type,
  type internal_messageservicev1_InternalMessageRecipient_Status as InternalMessageRecipient_Status,
  type internal_messageservicev1_SendMessageRequest as SendMessageRequest,
} from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useInternalMessageStore = defineStore('internal_message', () => {
  const internalMessageService = createInternalMessageServiceClient(
    requestClientRequestHandler,
  );

  const internalMessageRecipientService =
    createInternalMessageRecipientServiceClient(requestClientRequestHandler);

  /**
   * 查询消息列表
   */
  async function listMessage(
    noPaging: boolean = false,
    page?: number,
    pageSize?: number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await internalMessageService.ListMessage({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy: orderBy ?? [],
      query: makeQueryString(formValues ?? null),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 获取消息
   */
  async function getMessage(id: number) {
    return await internalMessageService.GetMessage({ id });
  }

  /**
   * 更新消息
   */
  async function updateMessage(id: number, values: object) {
    return await internalMessageService.UpdateMessage({
      id,
      data: {
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除消息
   */
  async function deleteMessage(id: number) {
    return await internalMessageService.DeleteMessage({
      id,
    });
  }

  /**
   * 获取用户的收件箱列表
   */
  async function listUserInbox(
    page?: number,
    pageSize?: number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    const noPaging: boolean = page === null || pageSize === null;
    return await internalMessageRecipientService.ListUserInbox({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy: orderBy ?? [],
      query: makeQueryString(formValues ?? null),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 将通知标记为已读
   */
  async function markNotificationAsRead(
    userId: number,
    recipientIds: number[],
  ) {
    return await internalMessageRecipientService.MarkNotificationAsRead({
      userId,
      recipientIds,
    });
  }

  /**
   * 删除收件箱中的通知
   */
  async function deleteNotificationFromInbox(
    userId: number,
    recipientIds: number[],
  ) {
    return await internalMessageRecipientService.DeleteNotificationFromInbox({
      userId,
      recipientIds,
    });
  }

  /**
   * 撤销某条消息
   */
  async function revokeMessage(userId: number, messageId: number) {
    return await internalMessageService.RevokeMessage({
      messageId,
      userId,
    });
  }

  /**
   * 发送消息
   */
  async function sendMessage(request: SendMessageRequest) {
    return await internalMessageService.SendMessage(request);
  }

  function $reset() {}

  return {
    $reset,
    listMessage,
    getMessage,
    updateMessage,
    deleteMessage,
    listUserInbox,
    sendMessage,
    revokeMessage,
    markNotificationAsRead,
    deleteNotificationFromInbox,
  };
});

export const internalMessageStatusList = computed(() => [
  {
    value: 'DRAFT',
    label: $t('enum.internalMessageStatus.DRAFT'),
  },
  {
    value: 'PUBLISHED',
    label: $t('enum.internalMessageStatus.PUBLISHED'),
  },
  {
    value: 'SCHEDULED',
    label: $t('enum.internalMessageStatus.SCHEDULED'),
  },
  {
    value: 'REVOKED',
    label: $t('enum.internalMessageStatus.REVOKED'),
  },
  {
    value: 'ARCHIVED',
    label: $t('enum.internalMessageStatus.ARCHIVED'),
  },
  {
    value: 'DELETED',
    label: $t('enum.internalMessageStatus.DELETED'),
  },
]);

export const internalMessageTypeList = computed(() => [
  {
    value: 'NOTIFICATION',
    label: $t('enum.internalMessageType.NOTIFICATION'),
  },
  {
    value: 'PRIVATE',
    label: $t('enum.internalMessageType.PRIVATE'),
  },
  {
    value: 'GROUP',
    label: $t('enum.internalMessageType.GROUP'),
  },
]);

export const internalMessageRecipientStatusList = computed(() => [
  {
    value: 'SENT',
    label: $t('enum.internalMessageRecipientStatus.SENT'),
  },
  {
    value: 'RECEIVED',
    label: $t('enum.internalMessageRecipientStatus.RECEIVED'),
  },
  {
    value: 'READ',
    label: $t('enum.internalMessageRecipientStatus.READ'),
  },
  {
    value: 'REVOKED',
    label: $t('enum.internalMessageRecipientStatus.REVOKED'),
  },
  {
    value: 'DELETED',
    label: $t('enum.internalMessageRecipientStatus.DELETED'),
  },
]);

export function internalMessageStatusLabel(
  value: InternalMessage_Status,
): string {
  const values = internalMessageStatusList.value;
  const matchedItem = values.find((item) => item.value === value);
  return matchedItem ? matchedItem.label : '';
}

export function internalMessageStatusColor(
  value: InternalMessage_Status,
): string {
  switch (value) {
    case 'ARCHIVED': {
      // 归档：已存档，用深灰色
      return '#6B7280';
    }
    case 'DELETED': {
      // 已删除：弱化显示，用浅灰色
      return '#E5E7EB';
    }
    case 'DRAFT': {
      // 草稿：未完成，用中灰色
      return '#9CA3AF';
    }
    case 'PUBLISHED': {
      // 已发布：成功状态，用绿色
      return '#10B981';
    }
    case 'REVOKED': {
      // 已撤回：异常状态，用红色
      return '#EF4444';
    }
    case 'SCHEDULED': {
      // 计划发送：待执行，用蓝色
      return '#3B82F6';
    }
    default: {
      // 新增未定义状态时，默认返回空（避免样式错误）
      return '';
    }
  }
}

export function internalMessageTypeLabel(value: InternalMessage_Type): string {
  const values = internalMessageTypeList.value;
  const matchedItem = values.find((item) => item.value === value);
  return matchedItem ? matchedItem.label : '';
}

export function internalMessageTypeColor(value: InternalMessage_Type): string {
  switch (value) {
    case 'GROUP': {
      // 群聊：多人互动，用活力感的颜色
      return '#10B981';
    } // 绿色（代表协作、活跃）
    case 'NOTIFICATION': {
      // 通知：系统/平台推送，用正式感的颜色
      return '#3B82F6';
    } // 蓝色（代表官方、提醒）
    case 'PRIVATE': {
      // 私信：一对一沟通，用私密感的颜色
      return '#8B5CF6';
    } // 紫色（代表个人、私密）
    default: {
      // 应对未定义类型，避免样式异常
      return '';
    }
  }
}

export function internalMessageRecipientStatusLabel(
  value: InternalMessageRecipient_Status,
): string {
  const values = internalMessageRecipientStatusList.value;
  const matchedItem = values.find((item) => item.value === value);
  return matchedItem ? matchedItem.label : '';
}

export function internalMessageRecipientStatusColor(
  value: InternalMessageRecipient_Status,
): string {
  switch (value) {
    case 'DELETED': {
      // 已删除：用户主动删除，视觉上弱化显示
      return '#E5E7EB';
    } // 浅灰色
    case 'READ': {
      // 已读：用户已查看，常规状态
      return '#6B7280';
    } // 深灰色
    case 'RECEIVED': {
      // 已接收（未读）：用户收到但未查看，需突出提醒
      return '#3B82F6';
    } // 蓝色（醒目，提示未读）
    case 'REVOKED': {
      // 已撤回：消息失效，带有异常含义
      return '#EF4444';
    } // 红色（警示，表明消息已失效）
    case 'SENT': {
      // 已发送（未接收）：消息发出但对方未确认接收，过渡状态
      return '#93C5FD';
    } // 浅蓝色（柔和，表示待接收）
    default: {
      // 应对未定义的状态，避免样式异常
      return '';
    }
  }
}
