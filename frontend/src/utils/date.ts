// 时间辅助函数
export function dateFormat(fmt, date) {
  let ret
  const opt = {
    'Y+': date.getFullYear().toString(), // 年
    'm+': (date.getMonth() + 1).toString(), // 月
    'd+': date.getDate().toString(), // 日
    'H+': date.getHours().toString(), // 时
    'M+': date.getMinutes().toString(), // 分
    'S+': date.getSeconds().toString() // 秒
    // 有其他格式化字符需求可以继续添加，必须转化成字符串
  }
  for (const k in opt) {
    ret = new RegExp('(' + k + ')').exec(fmt)
    if (ret) {
      fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, '0')))
    }
  }
  return fmt
}

export const pickerOptions = {
  'shortcuts': [
    {
      text: '今天',
      value: () => {

        let date = [dayjs().format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]

        return date
      },
    },
    {
      text: '昨天',
      value: () => {
        let date = [dayjs().subtract(-1, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().subtract(-1, 'day').format('YYYY-MM-DD 23:59:59')]


        return date
      },
    },
    {
      text: '一周',
      value: () => {
        let date = [dayjs().subtract(-7, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      },
    },
    {
      text: '两周',
      value: () => {
        let date = [dayjs().subtract(-14, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      },
    },
    {
      text: '一月',
      value: () => {
        let date = [dayjs().subtract(-30, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      },
    },
    {
      text: '三个月',
      value: () => {
        let date = [dayjs().subtract(-90, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      }
    },
    {
      text: '半年',
      value: () => {
        let date = [dayjs().subtract(-180, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      },
    },
    {
      text: '一年',
      value: () => {
        let date = [dayjs().subtract(-360, 'day').format('YYYY-MM-DD 00:00:00'), dayjs().format('YYYY-MM-DD 23:59:59')]


        return date
      },

    },
  ]
}
