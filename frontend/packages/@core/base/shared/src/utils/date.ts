import dayjs from 'dayjs';

/**
 * Format a date/time value into a specified string format.
 * @param time - The date/time value to format (number or string).
 * @param format - The desired output format (default is 'YYYY-MM-DD').
 * @returns The formatted date/time string.
 */
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

/**
 * Format a date/time value into 'YYYY-MM-DD HH:mm:ss' format.
 * @param time - The date/time value to format (number or string).
 * @returns The formatted date/time string.
 */
export function formatDateTime(time: number | string) {
  if (time === null || time === undefined || time === '') {
    return '';
  }
  return formatDate(time, 'YYYY-MM-DD HH:mm:ss');
}

/**
 * Check if a value is a Date object.
 * @param value - The value to check.
 * @returns True if the value is a Date object, false otherwise.
 */
export function isDate(value: any): value is Date {
  return value instanceof Date;
}

/**
 * Check if a value is a Dayjs object.
 * @param value - The value to check.
 * @returns True if the value is a Dayjs object, false otherwise.
 */
export function isDayjsObject(value: any): value is dayjs.Dayjs {
  return dayjs.isDayjs(value);
}
