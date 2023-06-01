<template>
  <div class="top-content">
    <div v-if="active == 0" class="addServer">
      <div class="notice">
        <span>提示：添加的服务器请确保可被SSH免密访问，并满足相关条件。</span>
      </div>
      <el-form
        :rules="rules"
        :model="newHost"
        class="server-form"
        ref="formRef"
        label-width="15%"
        labelPosition="right"
      >
        <el-form-item prop="ip" label="服务器IP：">
          <div class="server-ip-block">
            <!-- <el-row justify="space-between" :gutter="30">
              <el-col :span="12">
                <div class="grid-content">服务器IP地址</div>
              </el-col>
              <el-col :span="12">
                <div class="grid-content">
                  <el-button
                    class="addBtn"
                    type="text"
                    :icon="Plus"
                    @click="addServer"
                    >添加</el-button
                  >
                </div>
              </el-col>
            </el-row> -->
            <div class="addColItems" v-if="serverList.length > 0">
              <div
                v-for="(item, index) in serverList"
                :key="item"
                class="rowBlock"
              >
                <el-row justify="space-between" :gutter="30">
                  <el-col :span="12">
                    <div class="ipInput">
                      <el-input v-model="newHost.ip[index]" />
                    </div>
                  </el-col>
                  <el-col :span="12">
                    <el-button
                      v-if="index == 0"
                      type="text"
                      :icon="Plus"
                      @click="addServer"
                      class="addBtn"
                      >添加</el-button
                    >
                    <el-button
                      v-if="index != 0"
                      type="text"
                      :icon="Delete"
                      class="delBtn"
                      size="large"
                      @click="removeServer(index)"
                      >删除</el-button
                    >
                  </el-col>
                </el-row>
              </div>
            </div>
          </div>
        </el-form-item>
        <el-form-item label="节点角色：">
          <el-checkbox-group
            v-model="newHost.node_role"
            class="roleCheckbox checkbox"
          >
            <el-checkbox
              v-for="item in nodeRoleList"
              :key="item.id"
              :label="item.value"
              :name="item.label"
              >{{ item.label }}</el-checkbox
            >
          </el-checkbox-group>
        </el-form-item>
      </el-form>
    </div>
    <div v-if="active == 1" class="envInit" style="width: 100%">
      <div class="notice">
        <span
          >提示：输入SSH用户名和密码允许服务自动配置免密功能，否则需要手动配置每台服务器互相免密
        </span>
      </div>
      <el-form
        :rules="rules"
        :model="newHost"
        ref="formRef"
        class="initForm"
        label-width="15%"
        label-position="right"
      >
        <el-form-item prop="ssh_port" label="SSH端口号：">
          <el-input v-model="newHost.ssh_port" placeholder="22" />
        </el-form-item>
        <el-form-item label="SSH用户名：">
          <el-input v-model="newHost.ssh_user" disabled />
        </el-form-item>
        <el-form-item label="SSH密码：">
          <el-input
            v-model="newHost.ssh_password"
            :disabled="newHost.is_no_pass"
            type="password"
            placeholder="密码仅用于配置免密，不会保存"
          />
        </el-form-item>
        <el-form-item label="免密登录：">
          <el-switch
            v-model="newHost.is_no_pass"
            inline-prompt
            active-text="是"
            inactive-text="否"
          ></el-switch>
        </el-form-item>
      </el-form>
    </div>
    <!-- <div v-if="active == 2" class="configNetwork">
      <el-form
        ref="formRef"
        :rules="rules"
        :model="newHost"
        label-width="15%"
        label-position="right"
      >
        <el-form-item prop="public_network" label="服务器IP：">
          <div class="configNetwork-part">
            <el-row justify="space-between" :gutter="30">
              <el-col :span="8">
                <div class="grid-content">管理地址</div>
              </el-col>
              <el-col :span="8">
                <div class="grid-content">Public网络</div>
              </el-col>
              <el-col :span="8">
                <div class="grid-content">Cluster网络</div>
              </el-col>
            </el-row>
            <div class="addColItems" v-if="serverList.length > 0">
              <div v-for="(item, index) in serverList" :key="index">
                <el-row justify="center" :gutter="30">
                  <el-col :span="8">
                    <div class="ipInput">{{ newHost.ip[index] }}</div>
                  </el-col>
                  <el-col :span="8">
                    <div
                      v-if="newHost.ip[index] != '' && index == 0"
                      class="networkInput"
                    >
                      <el-input v-model="newHost.public_network" />
                    </div>
                    <div
                      v-if="newHost.ip[index] != '' && index != 0"
                      class="networkStr"
                    >
                      {{ newHost.public_network }}
                    </div>
                  </el-col>
                  <el-col :span="8">
                    <div
                      v-if="newHost.ip[index] != '' && index == 0"
                      class="networkInput"
                    >
                      <el-input v-model="newHost.cluster_network" />
                    </div>
                    <div
                      v-if="newHost.ip[index] != '' && index != 0"
                      class="networkStr"
                    >
                      {{ newHost.cluster_network }}
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
          </div> -->
    <!-- </el-form-item> -->
    <!-- </el-form> -->
    <!-- </div> -->
    <div v-if="active == 2" ref="formRef" class="instCheck">
      <!-- <el-form
        ref="formRef"
        :model="newHost"
        :rules="rules"
        label-width="15%"
        label-position="right"
      > -->
      <!-- <div class="check-rule"> -->
      <!-- <el-form-item label="" prop="content"> -->
      <div class="check-rule-content">
        <div
          v-for="(ruleItem, ruleIdx) in checkrules"
          :key="ruleIdx"
          class="check-rule-item"
        >
          <el-row :gutter="24">
            <el-col :span="10">
              <div class="check-rule-item-name">
                <b>{{ ruleItem.name }}：</b>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="check-rule-item-desc">
                {{ ruleItem.description }}
              </div>
            </el-col>
            <el-col :span="2">
              <div class="check-rule-item-status">
                <el-icon
                  class="check-rule-item-status-icon__success"
                  :size="20"
                  v-if="ruleItem.status == 'OK'"
                >
                  <circle-check-filled color="#67c23a" />
                </el-icon>
                <el-popover
                  v-if="ruleItem.status == 'FAILED'"
                  placement="right"
                  :width="600"
                  trigger="click"
                >
                  <p>{{ ruleItem.message }}</p>
                  <template #reference>
                    <el-icon class="check-rule-item-status-icon" :size="20">
                      <circle-close-filled color="#f56c6c" />
                    </el-icon>
                  </template>
                </el-popover>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
      <!-- </el-form-item> -->
      <!-- </div>
      </el-form> -->
    </div>
    <div v-if="active == 3" class="confirm">
      <el-form label-width="15%" label-position="right">
        <el-form-item prop="public_network" label="服务器IP：">
          <div class="confirm-part">
            <!-- <el-row justify="space-between" :gutter="30">
              <el-col :span="8">
                <div class="grid-content">管理地址</div>
              </el-col>
              <el-col :span="8">
                <div class="grid-content">Public网络</div>
              </el-col>
              <el-col :span="8">
                <div class="grid-content">Cluster网络</div>
              </el-col>
            </el-row> -->
            <div class="confirmView" v-if="serverList.length > 0">
              <div v-for="(item, index) in serverList" :key="index">
                <el-row justify="center" :gutter="30">
                  <el-col :span="24">
                    <div>{{ newHost.ip[index] }}</div>
                  </el-col>
                  <!-- <el-col :span="8">
                    <div class="networkInput">
                      <div class="networks">{{ newHost.public_network }}</div>
                    </div>
                  </el-col>
                  <el-col :span="8">
                    <div class="networkInput">
                      <div class="networks">{{ newHost.cluster_network }}</div>
                    </div>
                  </el-col> -->
                </el-row>
              </div>
            </div>
          </div>
        </el-form-item>
        <!-- <el-form-item label="类型：">
          {{ roleFormatter(newHost.node_type, 'type') }}
        </el-form-item> -->
        <el-form-item label="角色：">
          {{ roleFormatter(newHost.node_role) }}
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, computed } from 'vue'
import { Delete, Plus } from '@element-plus/icons-vue'
import { IServer, rule } from '@/service/host/types'
import { useStore } from '@/store'
import { rules } from '../../config/index-config'
import { CircleCheckFilled, CircleCloseFilled } from '@element-plus/icons-vue'

export default defineComponent({
  name: 'HostAddSteps',
  components: {
    CircleCheckFilled,
    CircleCloseFilled
  },
  setup() {
    const active = ref(0)
    const store = useStore()
    const formRef = ref()
    const checkrules = computed<rule[]>(
      () => store.getters['addServer/getRuleList']
    )

    let newServers: IServer = {
      ip: '',
      node_role: '',
      ssh_user: '',
      ssh_password: '',
      is_no_pass: false,
      ssh_port: '',
      public_network: '',
      cluster_network: ''
    }
    const nodeRoleList = ref([
      {
        id: 1,
        label: '管理节点',
        value: 'mon'
      },
      {
        id: 2,
        label: '存储节点',
        value: 'osd'
      },
      {
        id: 3,
        label: '对象存储网关',
        value: 'rgw'
      }
    ])

    const roleFormatter = computed(() => (data: any) => {
      let out = ''
      data.forEach((item: string) => {
        nodeRoleList.value.forEach((element) => {
          if (item == element.value) {
            out = out + element.label + '，'
          }
        })
      })
      out = out.substring(0, out.lastIndexOf('，'))
      return out
    })

    const hostInformationInit = () => {
      return {
        ip: [''],
        node_role: [],
        ssh_user: 'root',
        ssh_password: '',
        ssh_port: '22',
        is_no_pass: false,
        public_network: '',
        cluster_network: ''
      }
    }
    const newHost = reactive(hostInformationInit())
    const serverList = ref([''])

    const refreshStoreData = () => {
      newServers.ip = newHost.ip.join(',')
      newServers.node_role = newHost.node_role.join(',')
      newServers.ssh_port = newHost.ssh_port
      newServers.ssh_user = newHost.ssh_user
      newServers.ssh_password = newHost.ssh_password
      newServers.public_network = newHost.public_network
      newServers.cluster_network = newHost.cluster_network
      newServers.is_no_pass = newHost.is_no_pass
      store.commit('addServer/changeHostInfo', newServers)
    }

    const handleConnectionCheck = () => {
      // console.log('handle connection check called.')
      return newHost
    }

    const handleStepChange = (value: number) => {
      let out = false
      if (active.value == 2) {
        active.value = active.value + 1
        return true
      }
      formRef.value?.validate((valid: any) => {
        if (valid) {
          active.value = value
          refreshStoreData()
          out = true
        }
      })
      return out
    }

    const handlePrevStepChange = (value: number) => {
      active.value = value
      refreshStoreData
    }

    const addServer = () => {
      serverList.value.push('')
      newHost.ip.push('')
    }

    const removeServer = (idx: number) => {
      serverList.value.splice(idx, 1)
      newHost.ip.splice(idx, 1)
    }

    return {
      active,
      nodeRoleList,
      rules,
      formRef,
      serverList,
      roleFormatter,
      checkrules,
      Delete,
      Plus,
      newHost,
      addServer,
      removeServer,
      handleStepChange,
      handlePrevStepChange,
      handleConnectionCheck
    }
  }
})
</script>

<style lang="less">
.top-content {
  width: 100%;
  height: 100%;
  .el-form-item__error {
    font-size: 15px;
  }
  .addServer {
    .server-form {
      margin-top: 30px;

      .el-form-item__error {
        left: 62px;
      }
      .delBtn {
        color: #ec7259;
      }

      .server-ip-block {
        width: 100%;
      }

      .ipInput {
        display: flex;
        justify-content: center;
        margin-bottom: 2px;
      }

      .delete-icon:hover {
        cursor: pointer;
      }

      .typeCheckbox {
        margin-left: 61px;
      }
      .roleCheckbox {
        margin-left: 61px;
      }
    }
  }

  .initForm {
    margin-top: 30px;
  }

  .configNetwork {
    margin-top: 30px;
  }
  .configNetwork-part {
    width: 100%;
    .networkInput {
      display: flex;
      justify-content: center;
    }
    .networkStr {
      text-align: left;
      margin-left: 53px;
    }
  }
}

.instCheck {
  width: 100%;
  padding-top: 13px;
  .el-form-item__content {
    width: 100%;
  }
  .check-rule-content {
    width: 100%;
  }
  .check-rule-item-name {
    text-align: right;
    font-size: 1.2rem;
    padding-top: 15px;
    padding-right: 15px;
  }
  .check-rule-item-desc {
    padding-top: 15px;
    text-align: left;
  }
  .check-rule-item-status {
    padding-top: 15px;

    .check-rule-item-status-icon:hover {
      cursor: pointer;
    }
  }
}

.confirm {
  margin-top: 30px;
  text-align: left;
}
.confirm-part {
  width: 100%;
  .networks {
    display: flex;
    justify-content: center;
  }
  .confirmView {
    text-align: left;
  }
}

.notice {
  margin-top: 13px;
  margin-left: 15%;
  width: 70%;
  display: flex;
  align-items: center;
  background-color: RGB(84, 170, 249);
  color: white;
  padding: 5px;
  border-radius: 6px;
}
</style>
