export const rules = {
  ip: [
    {
      type: 'array',
      defaultField: {
        required: true,
        pattern:
          /^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$/,
        message: '请输入正确的IP地址',
        trigger: 'change'
      }
    }
  ],
  public_network: [
    {
      required: true,
      message: '请输入Public网络地址',
      trigger: 'blur'
    }
  ],
  ssh_port: [
    {
      required: true,
      message: '请输入SSH端口号',
      trigger: 'blur'
    },
    {
      pattern: /^\d+$/,
      message: 'SSH端口号必须为数字',
      trigger: 'blur'
    }
  ]
}
