import dayjs from 'dayjs';

export function formatDate(time: number | string, format = 'YYYY-MM-DD') {
  if (time === null || time === undefined || time === '') {
    return '';
  }
  if (isDate(time)) {
    return dayjs(time).format(format);
  }

  try {
    const date = dayjs(time);
    if (!date.isValid()) {
      throw new Error('Invalid date');
    }
    return date.format(format);
  } catch (error) {
    console.error(`Error formatting date: ${error}`);
    return time;
  }
}

export function formatDateTime(time: number | string) {
  if (time === null || time === undefined || time === '') {
    return '';
  }
  return formatDate(time, 'YYYY-MM-DD HH:mm:ss');
}

export function isDate(value: any): value is Date {
  return value instanceof Date;
}

export function isDayjsObject(value: any): value is dayjs.Dayjs {
  return dayjs.isDayjs(value);
}
