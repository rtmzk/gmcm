import { App } from 'vue'

import {
  ElAlert,
  ElButton,
  ElForm,
  ElInput,
  ElRadio,
  ElCheckbox,
  ElIcon,
  ElTabs,
  ElLink,
  ElContainer,
  ElMenu,
  ElDropdown,
  ElSteps,
  ElTable,
  ElDialog,
  ElSelect,
  ElDatePicker,
  ElRow,
  ElMessageBox,
  ElMessage,
  ElCol,
  ElCollapse,
  ElSwitch,
  ElScrollbar,
  ElResult,
  ElPopover,
  ElTooltip
} from 'element-plus'

const components = [
  ElButton,
  ElAlert,
  ElForm,
  ElInput,
  ElRadio,
  ElCheckbox,
  ElIcon,
  ElTabs,
  ElLink,
  ElContainer,
  ElMenu,
  ElDropdown,
  ElSteps,
  ElTable,
  ElDialog,
  ElSelect,
  ElRow,
  ElDatePicker,
  ElMessageBox,
  ElMessage,
  ElCol,
  ElCollapse,
  ElSwitch,
  ElScrollbar,
  ElResult,
  ElPopover,
  ElTooltip
]

export default function (app: App): void {
  for (const component of components) {
    app.use(component)
  }
}
