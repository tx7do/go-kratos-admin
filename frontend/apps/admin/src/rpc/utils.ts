const removeNullUndefined = (obj: any) =>
  Object.fromEntries(
    Object.entries(obj).filter(
      ([_, v]) => v !== null && v !== undefined && v !== '',
    ),
  );

/**
 * 创建查询字符串
 * @param formValues
 */
export function makeQueryString(formValues: null | object): null | string {
  if (formValues === null) {
    return null;
  }

  // 去除掉空值
  formValues = removeNullUndefined(formValues);

  // 过滤掉空对象
  if (Object.keys(formValues).length === 0) {
    return null;
  }

  // 简单的序列化成json字符串
  return JSON.stringify(formValues);
}

export function makeUpdateMask(keys: string[]): string {
  if (keys.length === 0) {
    return '';
  }
  if (!keys.includes('id')) {
    keys.push('id');
  }
  return keys.join(',');
}
