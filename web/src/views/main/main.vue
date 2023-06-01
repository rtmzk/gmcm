<template>
  <div class="main">
    <el-container class="main-content">
      <el-container class="page">
        <el-header class="page-header">
          <nav-header />
        </el-header>
        <el-main class="page-content">
          <div class="page-info">
            <router-view></router-view>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import NavHeader from '@/components/nav-header/src/nav-header.vue'
export default defineComponent({
  name: 'App',
  components: {
    NavHeader
  },
  setup() {
    return {}
  }
})
</script>

<style lang="less">
.main {
  position: flex;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.main-content,
.page {
  height: 100%;

  .page-info {
    margin: 0 7%;
  }
}

.page-content {
  height: calac(100% - 48px);
}

.el-header,
.el-footer {
  display: flex;
  color: #333;
  text-align: center;
}

.el-header {
  height: 59px !important;
  border-bottom: 1px solid RGB(231, 231, 231);
}

.el-main {
  color: #333;
  text-align: center;
  background-color: #f0f2f5;
}
</style>
