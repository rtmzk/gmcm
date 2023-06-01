<template>
  <div class="login">
    <div class="login-pannel">
      <p class="title">存储安装系统</p>
      <el-form
        :rules="rules"
        :model="accountLogin"
        class="loginForm"
        label-width="23%"
        ref="formRef"
      >
        <el-form-item class="formItem" label="账号:" prop="username">
          <el-input class="login-form-input" v-model="accountLogin.username" />
        </el-form-item>
        <el-form-item class="formItem" label="密码:" prop="password">
          <el-input
            class="login-form-input"
            v-model="accountLogin.password"
            show-password
            @keyup.enter="loginAction"
          />
        </el-form-item>
      </el-form>
      <div class="loginBtn">
        <el-button class="login-button" type="primary" @click="loginAction"
          >登录</el-button
        >
      </div>
    </div>

    <footer>
      <a href="https://x"
        >Copyright © 2007-2021 x</a
      >
    </footer>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { useStore } from 'vuex'
import { rules } from './config/account-config'

export default defineComponent({
  setup() {
    const store = useStore()
    const accountLogin = reactive({
      username: '',
      password: ''
    })
    const formRef = ref()

    const loginAction = () => {
      formRef.value?.validate((valid: any) => {
        if (valid) {
          store.dispatch('login/accountLoginAction', { ...accountLogin })
        }
      })
    }

    return {
      accountLogin,
      rules,
      formRef,
      loginAction
    }
  }
})
</script>

<style lang="less" scoped>
.login {
  width: 100%;
  height: 100%;
  background-image: url('../../assets/image/login_bg.png');
  display: flex;
  justify-content: center;
  align-items: center;

  .login-pannel {
    width: 600px;
    height: 400px;
    background-color: RGBA(137, 144, 159);
    border-radius: 10px;
    box-shadow: 0 0 30px RGBA(137, 144, 159, 0.6);

    .title {
      text-align: center;
      font-family: 'Microsoft YaHei', arial;
      font-size: 1.5rem;
      color: white;
      padding-top: 13px;
    }
    .loginForm {
      padding-top: 60px;
      .formItem {
        :deep(.el-form-item__label) {
          color: white;
          font-size: 18px;
          font-family: 'Microsoft YaHei', arial;
        }
        :deep(.el-form-item__label:before) {
          display: none;
        }
        :deep(.el-form-item__error) {
          font-size: 15px;
        }
      }
      .login-form-input {
        :deep(.el-input__inner) {
          border: 0 none;
          border-bottom: 1px solid #ccc;
          border-radius: 0px;
          padding-left: 5px;
          width: 77%;
          color: white;
          background-color: RGBA(137, 144, 159);
        }
        :deep(.el-input__inner:focus) {
          border-bottom: 1px solid RGB(64, 160, 255);
        }
        :deep(.el-input__suffix) {
          right: 22%;
        }
      }
    }
    .loginBtn {
      width: 100%;
      text-align: center;
      padding-top: 30px;
      :deep(.el-button--primary) {
        width: 40%;
      }
    }
  }

  footer {
    clear: both;
    display: block;
    text-align: center;
    margin: 0px auto;
    position: absolute;
    bottom: 50px;
    width: 100%;

    a {
      color: black;
      text-decoration: none;
      font-family: 'Microsoft YaHei', arial;
    }

    a:hover {
      transition: all 0.2s ease-in-out;
      color: RGB(64, 160, 255);
      font-family: 'Microsoft YaHei', arial;
    }
  }
}
</style>
