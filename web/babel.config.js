module.exports = {
  plugins: [
    [
      'import',
      {
        libraryName: 'element-plus',
        customName: (name) => {
          name = name.slice(3)
          return `element-plus/lib/components/${name}`
        },
        customStyleName: (name) => {
          name = name.slice(3)
          return `element-plus/lib/components/${name}/style/css`
        }
      }
    ]
  ],
  presets: ['@vue/cli-plugin-babel/preset']
}
