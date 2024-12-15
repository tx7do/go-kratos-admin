/**
 * 创建查询字符串
 * @param formValues
 */
export function makeQueryString(formValues: any): null | string {
  if (formValues === null) {
    return null;
  }

  // 去除掉空值
  for (const item in formValues) {
    if (
      formValues[item] === undefined ||
      formValues[item] === null ||
      formValues[item] === ''
    ) {
      // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
      delete formValues[item];
    }
  }

  // 过滤掉空对象
  if (Object.keys(formValues).length === 0) {
    return null;
  }

  // 简单的序列化成json字符串
  return JSON.stringify(formValues);
}

export function makeUpdateMask(keys: string[]): string {
  keys.push('id');
  return keys.join(',');
}

export const statusList = [
  { value: 'ON', label: '正常' },
  { value: 'OFF', label: '停用' },
];
