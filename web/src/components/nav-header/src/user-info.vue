<template>
  <el-dropdown class="el-dropdown-link" trigger="click">
    <el-button type="text">
      {{ name }}
      <el-icon :size="12"><arrow-down /></el-icon>
    </el-button>
    <template #dropdown>
      <el-dropdown-menu>
        <!-- <el-dropdown-item>用户信息</el-dropdown-item>
        <el-dropdown-item>修改密码</el-dropdown-item> -->
        <el-dropdown-item @click="logoutHandler">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useStore } from '@/store'
import { accountLogoutRequest } from '@/service/login/login'
import localCache from '@/utils/cache'
import { useRouter } from 'vue-router'
import { ArrowDown } from '@element-plus/icons-vue'
export default defineComponent({
  components: {
    ArrowDown
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const name = computed(() => store.state.login.username)
    const logoutHandler = () => {
      accountLogoutRequest().then(() => {
        localCache.deleteCache('token')
        router.push('/login')
      })
    }
    return {
      name,
      logoutHandler
    }
  }
})
</script>

<style lang="less" scoped>
.el-dropdown-link {
  width: 80px;
}
</style>
