<template>
  <div class="steps-content">
    <el-steps
      :active="active"
      finish-status="finish"
      simple
      style="margin-top: 20px"
    >
      <el-step title="添加服务器"> </el-step>
      <el-step title="环境初始化"></el-step>
      <!-- <el-step title="配置网络"></el-step> -->
      <el-step title="安装检查"></el-step>
      <el-step title="确认信息"></el-step>
    </el-steps>
    <div>
      <host-add-steps ref="sonRef" />
    </div>
    <div class="footer">
      <hr />
      <div class="prevstep">
        <el-button @click="handlePrevAction" v-if="active > 0"
          >上一步
        </el-button>
      </div>

      <div class="button">
        <el-button
          v-if="active == 1"
          type="primary"
          @click="handleConnectionCheck"
          >连接测试</el-button
        >
        <el-button v-if="active == 2" type="primary" @click="handleEnvInit"
          >环境检测</el-button
        >

        <el-button
          v-if="active != 3"
          :disabled="disabled"
          class="nextStep"
          @click="handleActiveChange"
          type="primary"
          >下一步</el-button
        >

        <el-button
          v-if="active == 3"
          :value="value"
          class="nextStep"
          type="primary"
          @click="handleServerAdd"
          >完成</el-button
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import 'element-plus/theme-chalk/el-step.css'
import HostAddSteps from './host-add-steps.vue'
import {
  connectionCheckRequest,
  hostAddRequest,
  hostEnvcActionRequest
} from '@/service/host/hosts'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
export default defineComponent({
  components: {
    HostAddSteps
  },
  props: {
    value: Boolean
  },
  setup() {
    const active = ref(0)
    const disabled = ref(false)
    const sonRef = ref()
    const store = useStore()
    const hostInfo = computed(() => store.getters['addServer/getHostInfo'])

    const handleActiveChange = () => {
      const nextStepOr = sonRef.value.handleStepChange(active.value + 1)
      // console.log(nextStepOr)
      if (nextStepOr) {
        active.value = active.value + 1
      }
      if (active.value == 1 || active.value == 2) {
        disabled.value = true
      }
    }
    const handlePrevAction = () => {
      active.value = active.value - 1
      if (active.value != 2) {
        disabled.value = false
      }

      sonRef.value.handlePrevStepChange(active.value)
    }

    const handleConnectionCheck = () => {
      // console.log('handle connection check.')
      const data = sonRef.value.handleConnectionCheck()
      connectionCheckRequest(data).then((res) => {
        // console.log(res)
        if (res.successes) {
          ElMessage({
            message: '连接成功',
            type: 'success'
          })
          disabled.value = false
        }
        if (res.message) {
          ElMessage({
            message: res.code + ': ' + res.message,
            type: 'error'
          })
        }
      })
    }

    const handleEnvInit = () => {
      let failedCount = 0
      hostEnvcActionRequest(hostInfo.value).then((res) => {
        if (res.rules) {
          store.commit('addServer/changeRuleInfo', res.rules)
          res.rules.forEach((e) => {
            if (e.status == 'FAILED') {
              failedCount = failedCount + 1
            }
          })
          if (failedCount == 0) {
            store.commit('addServer/changeCheckStatus', true)
            disabled.value = false
          }
        }
      })
    }

    const handleServerAdd = () => {
      hostAddRequest(store.getters['addServer/getHostInfo']).then((res) => {
        if (res.successes == 'ok') {
          ElMessage.success({
            type: 'success',
            message: '添加成功'
          })

          setTimeout(() => {
            store.commit('addServer/changeOpenStatus', false)
            location.reload()
          }, 1000)
        }
      })
    }
    return {
      active,
      sonRef,
      disabled,
      handleActiveChange,
      handlePrevAction,
      handleEnvInit,
      handleServerAdd,
      handleConnectionCheck
    }
  }
})
</script>

<style lang="less">
.footer {
  display: flex;
  text-align: right;
  justify-content: space-between;
  margin: 0px auto;
  width: 100%;

  .nextStep {
    margin: 10px;
  }

  .prevstep {
    padding: 10px;
  }
}
</style>
