<template>
  <div class="nav-menu">
    <div class="logo">
      <img class="img" src="@/assets/logo.png" alt="" />
      <!-- <span v-if="!collapse" class="title">Manage</span> -->
    </div>
    <el-menu
      class="el-memu-vertical"
      background-color="#001529"
      text-color="#fff"
      :collapse="collapse"
    >
      <el-menu-item
        :index="HostListPageLocation.id"
        @click="handleMenuItemClick(HostListPageLocation)"
      >
        <span>{{ HostListPageLocation.name }}</span>
      </el-menu-item>
      <!-- <el-menu-item
        :index="ClusterInitPageLocation.id"
        @click="handleMenuItemClick(ClusterInitPageLocation)"
      >
        <span>{{ ClusterInitPageLocation.name }}</span>
      </el-menu-item> -->
    </el-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
export default defineComponent({
  props: {
    collapse: {
      type: Boolean,
      default: false
    }
  },
  setup() {
    const HostListPageLocation = {
      url: '/index/hosts',
      id: '1',
      icon: '',
      name: '主机列表'
    }
    // const ClusterInitPageLocation = {
    //   url: '/index/cluster',
    //   id: '2',
    //   icon: '',
    //   name: '集群初始化'
    // }
    const isCollapse = ref(false)
    const router = useRouter()
    const handleMenuItemClick = (item: any) => {
      router.push({
        path: item.url
      })
    }
    return {
      isCollapse,
      HostListPageLocation,
      // ClusterInitPageLocation,
      handleMenuItemClick
    }
  }
})
</script>

<style lang="less" scoped>
.nav-menu {
  height: 100%;
  background-color: #001529;

  .logo {
    display: flex;
    height: 28px;
    padding: 12px 10px 8px 10px;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;

    .img {
      height: 100%;
      margin: 0 10px;
    }

    .title {
      font-size: 16px;
      font-weight: 700;
      color: white;
    }
  }

  .el-menu {
    border-right: none;
  }
  // 目录
  .el-submenu {
    background-color: #001529 !important;
    // 二级菜单 ( 默认背景 )
    .el-menu-item {
      padding-left: 50px !important;
      background-color: #0c2135 !important;
    }
  }

  ::v-deep .el-submenu__title {
    background-color: #001529 !important;
  }

  // hover 高亮
  .el-menu-item:hover {
    color: #fff !important; // 菜单
  }

  .el-menu-item.is-active {
    color: #fff !important;
    background-color: #5194d6 !important;
  }
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 100%;
  height: calc(100% - 48px);
}
</style>
