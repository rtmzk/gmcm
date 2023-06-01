<template>
  <div class="nav-header">
    <div class="logo">
      <img class="top-logo" src="@/assets/image/header-top-logo.svg" />
    </div>
    <div class="content">
      <!-- <div class="menu">
        <el-menu
          mode="horizontal"
          text-color="RGB(81, 90, 110)"
          active-text-color="RGB(51, 153, 255)"
        >
          <el-menu-item
            class="menu-item"
            :index="mainPage.id"
            @click="handleMenuItemClick(mainPage)"
            style="padding: 0 35px"
            ><h3>{{ mainPage.name }}</h3>
          </el-menu-item>
          <el-menu-item
            class="menu-item"
            :index="managePage.id"
            @click="handleMenuItemClick(managePage)"
            style="padding: 0 35px"
            ><h3>{{ managePage.name }}</h3>
          </el-menu-item>
        </el-menu>
      </div> -->
      <user-info />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import UserInfo from './user-info.vue'
import { useRouter } from 'vue-router'
export default defineComponent({
  components: {
    UserInfo
  },
  setup() {
    // const mainPage = {
    //   id: '1',
    //   name: '首页',
    //   url: '/main/index'
    // }
    // const managePage = {
    //   id: '2',
    //   name: '系统管理',
    //   url: '/main/system'
    // }
    const router = useRouter()
    const handleMenuItemClick = (item: any) => {
      router.push({
        path: item.url
      })
    }
    return {
      // mainPage,
      // managePage,
      handleMenuItemClick
    }
  }
})
</script>

<style lang="less">
.nav-header {
  width: 100%;
  display: flex;
  align-items: center;

  .top-logo {
    width: 180px;
    padding-top: 7px;
  }
  .top-logo:hover {
    cursor: pointer;
  }
  .content {
    display: block;
    position: absolute;
    right: 10px;
  }
  // .content {
  //   width: calc(100% -200px);
  //   height: 100%;
  //   display: flex;
  //   justify-content: space-between;
  //   align-items: center;
  //   flex: 1;
  //   padding: 0 20px;
  //   .menu {
  //     padding-left: 40px;
  //     width: calc(100% - 200px);
  //     height: 100%;

  //     h3 {
  //       margin: 0;
  //     }
  //   }
  // }
}
</style>
