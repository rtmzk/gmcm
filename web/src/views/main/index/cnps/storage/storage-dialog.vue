<template>
  <div class="top">
    <div v-if="devs.devices == null">暂无可用节点或存储</div>
    <div v-if="devs.devices != null" class="content">
      <el-form
        :rules="{
          replicas: [
            {
              type: 'integer',
              required: true,
              min: 1,
              max: devs.devices.length,
              message: '副本数必须是数字且最小为1,最多不超过节点数量',
              trigger: 'blur'
            }
          ],
          public_network: [
            {
              required: true,
              message: '请输入Public网络地址',
              trigger: 'blur'
            }
          ]
        }"
        ref="formRef"
        :model="devs"
        label-width="15%"
      >
        <el-form-item prop="replicas" label="副本数：">
          <el-input type="number" v-model.number="devs.replicas" />
        </el-form-item>

        <el-form-item prop="public_network" label="Public网络：">
          <el-input v-model="devs.public_network" />
          <el-tooltip
            effect="dark"
            content="<span>Public网络为存储集群对外网络,客户端可以通过该网络访问到存储池</span>"
            placement="right"
            raw-content
          >
            <el-icon :size="16" style="padding-left: 5px">
              <question-filled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="Cluster网络：">
          <el-input v-model="devs.cluster_network" />
          <el-tooltip
            effect="dark"
            content="<span>Cluster网络为集群内部网络，用于存储集群各个节点之间副本数据同步，默认和Public网络使用同一个网络</span>"
            placement="right"
            raw-content
          >
            <el-icon :size="16" style="padding-left: 5px">
              <question-filled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="磁盘开启总数：">{{
          calculateTotalEnableCount(devs)
        }}</el-form-item>
        <el-form-item label="磁盘设置：">
          <div class="diskBlock">
            <div class="diskHeader">
              <el-row :gutter="48">
                <el-col :span="6">
                  <div class="grid-content bg-purple">主机列表</div>
                </el-col>
                <el-col :span="6">
                  <div class="grid-content bg-purple">开启磁盘数量</div>
                </el-col>
                <el-col :span="6">
                  <div class="grid-content bg-purple">大小</div>
                </el-col>
              </el-row>
            </div>
            <div class="diskContent">
              <div class="disks">
                <el-collapse>
                  <el-collapse-item
                    v-for="(item, index) in devs.devices"
                    :key="index"
                    :name="index"
                  >
                    <template #title>
                      <el-row class="content-row" :gutter="48">
                        <el-col :span="6">
                          <div class="grid-content bg-purple">
                            {{ item.ip }}
                          </div>
                        </el-col>
                        <el-col class="disk-number" :span="6">
                          <div class="grid-content bg-purple disk-num">
                            {{ calculateEnableCount(item.device) }}
                          </div>
                        </el-col>
                        <el-col class="total-disk-size" :span="10">
                          <div class="grid-content bg-purple">
                            {{ caculateDeviceSize(item.device) }}
                          </div>
                        </el-col>
                      </el-row>
                    </template>
                    <div class="drop-list">
                      <div class="drop-list-header">
                        <el-row :gutter="20">
                          <el-col :span="4">名称</el-col>
                          <el-col :span="4">磁盘类型</el-col>
                          <el-col :span="4">磁盘大小</el-col>
                          <el-col :span="4">禁用/启用</el-col>
                          <el-col :span="4">缓存设备</el-col>
                        </el-row>
                      </div>
                      <div
                        v-for="(v, k) in item.device"
                        :key="k"
                        class="drop-list-content"
                      >
                        <el-row :gutter="20">
                          <el-col :span="4">{{ v.name }}</el-col>
                          <el-col :span="4">
                            <el-select
                              size="small"
                              v-model="v.type"
                              :placeholder="v.type"
                            >
                              <el-option
                                v-for="z in option"
                                :key="z.value"
                                :label="z.label"
                                :value="z.value"
                              ></el-option>
                            </el-select>
                          </el-col>
                          <el-col :span="4">{{ v.size }}</el-col>
                          <el-col :span="4">
                            <el-switch
                              size="small"
                              v-model="v.enabled"
                            ></el-switch>
                          </el-col>
                          <el-col :span="4">
                            <el-switch
                              v-if="v.type === 'ssd'"
                              size="small"
                              v-model="v.cached"
                            ></el-switch>
                          </el-col>
                        </el-row>
                      </div>
                    </div>
                  </el-collapse-item>
                </el-collapse>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
  <footer class="footer">
    <div class="operation">
      <el-button @click="cancelInitStorage">取消</el-button>
      <el-button
        type="primary"
        :disabled="
          devs.devices == null || devs.devices.length == 0 ? true : false
        "
        @click="initStorageCommit"
        >初始化</el-button
      >
    </div>
  </footer>
</template>
<script lang="ts">
import { StorageInspect, Device } from '@/store/main/servers/types'
import { defineComponent, ref, computed } from 'vue'
import { useStore } from 'vuex'
import { QuestionFilled } from '@element-plus/icons-vue'
export default defineComponent({
  name: 'StorageDialog',
  components: {
    QuestionFilled
  },
  setup() {
    const store = useStore()
    const devs = computed<StorageInspect>(
      () => store.getters['server/getDevices']
    )
    const maxReplicas = ref<number>(1)
    const formRef = ref()
    const hosts = ref(store.getters['server/getHosts'])

    if (devs.value.devices != null && devs.value.devices.length > 0) {
      maxReplicas.value = devs.value.devices.length
      devs.value.public_network =
        devs.value.devices[0].ip.split('.', 3).join('.') + '.0/24'
    }

    let Status = {
      status: 1
    }

    const calculateEnableCount = computed(() => (data: Device[]) => {
      let enableCount = 0
      for (let i = 0; i < data.length; i++) {
        if (data[i].enabled) {
          enableCount = enableCount + 1
        }
      }
      return enableCount
    })

    const calculateTotalEnableCount = computed(() => (data: StorageInspect) => {
      let enableCount = 0
      for (let i = 0; i < data.devices.length; i++) {
        for (let j = 0; j < data.devices[i].device.length; j++) {
          if (data.devices[i].device[j].enabled) {
            enableCount += 1
          }
        }
      }
      return enableCount
    })

    const caculateDeviceSize = computed(() => (data: Device[]) => {
      let sizeMap = new Map()
      let out = ''
      for (let i = 0; i < data.length; i++) {
        if (data[i].enabled) {
          if (sizeMap.get(data[i].size) != undefined) {
            sizeMap.set(data[i].size, sizeMap.get(data[i].size) + 1)
          } else {
            sizeMap = sizeMap.set(data[i].size, 1)
          }
        }
      }

      sizeMap.forEach((value, key) => {
        out = out + key + '*' + value + '，'
      })

      out = out.substring(0, out.lastIndexOf('，'))
      return out
    })

    const option = [
      {
        value: 'hdd',
        label: 'HDD'
      },
      {
        value: 'ssd',
        label: 'SSD'
      }
    ]
    const cancelInitStorage = () => {
      store.commit('initStorage/changeOpenStatus', false)
    }

    const initStorageCommit = () => {
      formRef.value?.validate((valid: any) => {
        if (valid) {
          store.dispatch(
            'server/initHostStorage',
            store.getters['server/getDevices']
          )
          store.dispatch('initStorage/setInitStatusReloadBtn', Status)
          store.commit('initStorage/changeOpenStatus', false)
        }
      })
    }

    return {
      devs,
      hosts,
      option,
      formRef,
      cancelInitStorage,
      initStorageCommit,
      calculateEnableCount,
      calculateTotalEnableCount,
      caculateDeviceSize,
      maxReplicas
    }
  }
})
</script>

<style lang="less">
.top {
  height: 100%;
  width: 100%;
  padding-bottom: 20px;
  .el-form-item__error {
    font-size: 15px;
  }
  .content {
    .el-input {
      width: 50%;
    }
    .diskBlock {
      width: 100%;

      .diskHeader {
        text-align: left;
      }

      .diskContent {
        .content-row {
          width: 100%;
          height: 100%;
          display: flex;
          justify-items: center;
          align-content: center;
          text-align: left;

          .disk-num {
            margin-left: 49px;
          }
          .total-disk-size {
            margin-left: 23px;
          }
        }
        .drop-list {
          .drop-list-header {
            background-color: RGB(229, 229, 229);
            text-align: left;
          }
          .drop-list-content {
            text-align: left;
            padding-top: 4px;
          }
        }
      }
    }
  }
}

.footer {
  height: 100%;
  .operation {
    display: block;
    position: absolute;
    bottom: 10px;
    right: 30px;
  }
}
</style>
