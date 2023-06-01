<template>
  <div class="node">
    <div class="node-header">
      <el-button @click="handleAddServer" type="primary" class="add-host-btn"
        >添加主机</el-button
      >
      <el-button
        :disabled="initDisabled"
        @click="handleInitStorage"
        type="primary"
        class="inst-btn"
        >初始化存储</el-button
      >
      <el-button @click="handleLogFetch" type="primary" class="log-btn"
        >查看安装日志
      </el-button>
      <div class="host-dialog">
        <el-dialog
          v-model="openDialog"
          title="添加主机"
          width="55%"
          top="25vh"
          destroy-on-close
          :before-close="handleClose"
        >
          <div class="steps">
            <host-dialog :value="openDialog" />
          </div>
        </el-dialog>
      </div>
      <div class="storage-dialog">
        <el-dialog
          v-model="openStorageDialog"
          title="初始化存储"
          width="55%"
          top="25vh"
          destroy-on-close
          :before-close="handleClose"
        >
          <div class="storageStep">
            <storage-dialog />
          </div>
        </el-dialog>
      </div>
      <div class="storage-install-logs">
        <el-dialog
          v-model="openStorageInstallLogDialog"
          title="安装日志"
          width="70%"
          destroy-on-close
          :before-close="handleCloseLogPannel"
        >
          <div class="log-pannel">
            <div class="scrollbar">
              {{ logData }}
            </div>
          </div>
        </el-dialog>
      </div>
    </div>
    <el-table
      :data="tableData"
      empty-text="暂无可用节点"
      style="width: 100%; margin-top: 20px"
    >
      <el-table-column prop="ip" label="主机ip" />
      <!-- <el-table-column prop="login_type" label="添加形式" width="200" /> -->
      <el-table-column prop="metadata.createdAt" label="添加时间" />
      <!-- <el-table-column prop="node_type" label="节点类型" /> -->
      <el-table-column prop="node_role" label="节点角色" />
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope" v-if="tableData.length > 0">
          <el-button
            class="deleteBtn"
            type="text"
            @click="handleHostDelete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import HostDialog from './cnps/hosts/host-dialog.vue'
import StorageDialog from './cnps/storage/storage-dialog.vue'
import {
  hostListRequest,
  hostDeleteRequest,
  hostInitLogFetchRequest
} from '@/service/host/hosts'
import { TimeFormatter } from '@/utils/time'
import { useStore } from 'vuex'
export default defineComponent({
  name: 'hosts',
  components: {
    HostDialog,
    StorageDialog
  },
  emits: ['tableDataChange'],
  setup() {
    const store = useStore()
    const Loading = ref(false)
    const disabled = ref(true)
    const initDisabled = computed(
      () => store.getters['initStorage/getBtnstatus']
    )
    // console.log(initDisabled.value)
    const openDialog = computed(() => store.getters['addServer/getOpenStatus'])
    const openStorageDialog = computed(
      () => store.getters['initStorage/getOpenStatus']
    )
    // const status = computed(() => store.getters['server/getStatus'])
    // console.log(status.value)
    const openStorageInstallLogDialog = computed(
      () => store.getters['initStorage/getOpenLogPannelStatus']
    )
    const tableData = ref([])
    const logData = ref('')
    const offset = ref(0)
    let timer: any

    const handleClose = (done: any) => {
      ElMessageBox.confirm('确定要退出吗?', {
        confirmButtonText: '确认',
        cancelButtonText: '取消'
      }).then(() => {
        store.commit('addServer/changeOpenStatus', false)
        store.commit('initStorage/changeOpenStatus', false)
        done()
      })
    }

    const tableDataRefreshClick = () => {
      hostListRequest().then((res) => {
        if (res.hosts) {
          res.hosts.forEach((item: any) => {
            item.metadata.createdAt = TimeFormatter(item.metadata.createdAt)
            item.metadata.updatdAt = TimeFormatter(item.metadata.updateAt)
          })
        }
        tableData.value = res.hosts
      })
      store.dispatch('server/serverList')
      store.dispatch('server/deviceList')
      store.dispatch('initStorage/loadBtnStatus')
    }

    const handleAddServer = () => {
      store.commit('addServer/changeOpenStatus', true)
    }

    const handleInitStorage = () => {
      store.commit('initStorage/changeOpenStatus', true)
    }

    const handleHostDelete = (row: any) => {
      ElMessageBox.confirm('确定要删除该节点吗?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        center: true,
        draggable: true,
        type: 'warning'
      })
        .then(() => {
          hostDeleteRequest(row.metadata.id).then((resp) => {
            if (resp.successes) {
              tableData.value.forEach((v: any, k: number) => {
                if (v.metadata.id == row.metadata.id) {
                  tableData.value.splice(k, 1)
                  store.commit('server/changeHosts', tableData.value)
                  store.dispatch('server/deviceList')
                }
              })
              ElMessage.success({
                type: 'successes',
                message: '删除成功'
              })
            } else {
              ElMessage.error('删除失败')
            }
          })
        })
        .catch(() => {
          ElMessage({
            type: 'info',
            message: '取消删除'
          })
        })
    }

    const handleLogFetch = () => {
      store.commit('initStorage/changeOpenLogPannelStatus', true)
      timer = setInterval(function () {
        hostInitLogFetchRequest(offset.value).then((res) => {
          if (offset.value == 0) {
            logData.value = ''
          }
          if (res.message) {
            logData.value = res.message
          }
          if (res.text) {
            offset.value = res.offset
            logData.value = logData.value + res.text
          }
        })
      }, 2000)
    }
    const handleCloseLogPannel = () => {
      clearInterval(timer)
      store.commit('initStorage/changeOpenLogPannelStatus', false)
    }
    // onUpdated(() => {
    //   nextTick(() => {
    //     let container = document.getElementsByClassName('log-pannel')[0]
    //     container.scrollTop = container.scrollHeight
    //   })
    // })

    tableDataRefreshClick()
    return {
      Loading,
      disabled,
      openDialog,
      openStorageDialog,
      openStorageInstallLogDialog,
      handleInitStorage,
      handleClose,
      tableDataRefreshClick,
      handleHostDelete,
      handleAddServer,
      handleLogFetch,
      handleCloseLogPannel,
      initDisabled,
      tableData,
      logData
    }
  }
})
</script>

<style lang="less">
.node {
  .node-header {
    display: flex;
    width: 100%;
    margin-top: 20px;

    .el-dialog {
      justify-items: center;
    }
  }

  .el-input {
    width: 70%;
    display: flex;
    justify-items: left;
  }

  .checkbox {
    display: flex;
    justify-items: left;
  }
  .deleteBtn:hover {
    color: #ec7259;
  }
}
.log-pannel {
  display: block;
  height: 600px;
  white-space: pre-line;
  text-align: left;
  overflow-y: auto;
  overflow-x: hidden;
  transition: transform 0.25s ease-out;
  z-index: 3;
}

.log-pannel::-webkit-scrollbar {
  width: 7px;
}

.log-pannel::-webkit-scrollbar-thumb {
  background: transparent;
  border-radius: 4px;
}

.log-pannel:hover::-webkit-scrollbar-thumb {
  background: hsla(0, 0%, 53%, 0.4);
}

.log-pannel:hover::-webkit-scrollbar-track {
  background: hsla(0, 0%, 53%, 0.1);
}
</style>
