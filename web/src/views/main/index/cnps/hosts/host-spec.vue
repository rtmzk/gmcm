<template>
  <el-table
    @tableDataChange="handleTableDataChange"
    :data="tableData"
    empty-text="暂无可用节点"
    style="width: 100%; margin-top: 20px"
  >
    <el-table-column prop="ip" label="主机ip" width="200" />
    <el-table-column prop="node_type" label="添加形式" width="200" />
    <el-table-column prop="metadata.createdAt" label="添加时间" />
    <el-table-column prop="node_type" label="节点类型" />
    <el-table-column prop="node_role" label="节点角色" />
    <el-table-column fixed="right" label="操作" width="200">
      <template #default="scope" v-if="tableData.length > 0">
        <!-- <el-button type="text" @click="handleViewClick">查看</el-button>
        <el-button type="text" @click="handleEditClick">编辑</el-button>
        <el-button type="text" @click="handleInitClick">初始化</el-button> -->
        <el-button
          class="deleteBtn"
          type="text"
          @click="handleHostDelete(scope.row)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { hostDeleteRequest, hostListRequest } from '@/service/host/hosts'
export default defineComponent({
  setup() {
    const tableData = ref([])
    // console.log(props)
    hostListRequest().then((res) => {
      if (res.hosts) {
        tableData.value = res.hosts
      }
    })
    // const handleViewClick = () => {
    //   console.log('detail function called.')
    // }
    // const handleEditClick = () => {
    //   console.log('edit function called.')
    // }
    // const handleInitClick = () => {
    //   console.log('init function called.')
    //   console.log(tableData)
    // }
    const handleHostDelete = (row: any) => {
      hostDeleteRequest(row.metadata.id).then(() => {
        tableData.value.forEach((v: any, k: number) => {
          if (v.metadata.id == row.metadata.id) {
            tableData.value.splice(k, 1)
          }
        })
      })
    }
    const handleTableDataChange = (data: any) => {
      tableData.value = data
    }

    return {
      // handleViewClick,
      // handleEditClick,
      // handleInitClick,
      handleHostDelete,
      handleTableDataChange,
      tableData
    }
  }
})
</script>

<style lang="less">
.deleteBtn:hover {
  color: #ec7259;
}
</style>
