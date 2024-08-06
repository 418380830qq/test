<template>
  <el-input v-model="searchKeyword" placeholder="输入关键词或关键词ID进行搜索" style="width: 200px; margin-bottom: 10px;">
  </el-input>
  <el-button @click="fetchData">获取当前的数据</el-button>
  <div id="app">
    <el-table :data="currentPageDatas" border style="width: 100%">
      <!-- 添加选择列 -->
      <el-table-column width="60">
        <template v-slot="scope">
          <el-switch v-model="scope.row.Status" @change="handleSwitchChange(scope.row)" />
        </template>
      </el-table-column>
      <!-- 其他列 -->
      <el-table-column prop="CreatedAt" label="创建时间"></el-table-column>
      <el-table-column prop="UpdatedAt" label="更新时间"></el-table-column>
      <el-table-column prop="Wordid" label="关键词ID" sortable></el-table-column>
      <el-table-column prop="PlanName" label="计划名称"></el-table-column>
      <el-table-column prop="Word" label="关键词"></el-table-column>
      <el-table-column prop="MinPrice" label="最低出价"></el-table-column>
      <el-table-column prop="Maxprice" label="top1出价"></el-table-column>
      <el-table-column prop="SuggestPriceStr" label="建议出价"></el-table-column>
      <el-table-column prop="SuggestPriceStr" label="追加价格">
        <template v-slot="scope">
          <el-input v-model="scope.row.UpBySuggest" @change="updateUpBySuggest(scope.row)"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="出价最高价格" width="120">
        <template v-slot="scope">
          <el-input v-model="scope.row.Setprice" @change="updateMaxprice(scope.row)"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="当前价格" width="120">
        <template v-slot="scope">
          <el-input v-model="scope.row.NowPrice" @change="updateNowprice(scope.row)"></el-input>
        </template>
      </el-table-column>
      <el-table-column prop="Remarks" label="备注"></el-table-column>
    </el-table>

    <!-- 分页器 -->
    <el-pagination :page-count="pageCount" :current-page.sync="currentPage" :prev-text="'上一页'" :next-text="'下一页'"
      @current-change="handlePageClick" />
  </div>



</template>


<script>
import { defineComponent } from 'vue';
import axios from 'axios';

export default defineComponent({
  data() {
    return {
      message: [], // 初始化为空数组
      currentPage: 1, // 当前页码
      itemsPerPage: 10, // 每页显示数量
      editable: false,
      productId: '', // 产品 ID
      newValue: '', // 新值
      filteredData: [], // 过滤后的数据
      searchKeyword: '', // 搜索关键词
      selectedRows: [],   // 存储被选中的行数据
    };
  },
  computed: {
    pageCount() {
      return Math.ceil(this.message.length / this.itemsPerPage); // 总页数
    },
    currentPageDatas() {
      // 根据当前页和每页显示条目数计算起始索引和结束索引
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;
      // 根据搜索关键词过滤数据
      if (this.searchKeyword.trim() === '') {
        // 如果关键词为空，直接从原始数据中获取当前页的数据
        this.filteredData = this.message.slice(startIndex, endIndex);
      } else {
        // 如果有搜索关键词，先过滤数据，然后获取当前页的数据
        const keyword = this.searchKeyword.trim().toLowerCase();
        this.filteredData = this.message.filter(item =>
          item.Word.toLowerCase().includes(keyword) || (item.Wordid.toString().includes(keyword))
        ).slice(startIndex, endIndex);
      }
      // 返回当前页的数据
      return this.filteredData;
    }
  },
  methods: {
    fetchData() {
      axios.post('/fetch')
        .then(response => {
          if (response.data.message === "执行成功") {
            let alertInstance = this.$alert('数据更新完成', '提示', {
              confirmButtonText: '确定',
              callback: action => {
                this.getData(); // 在确认按钮点击后获取数据
              }
            });

            // 设置定时器，在1分钟后关闭提示框
            setTimeout(() => {
              this.getData();
              if (alertInstance.value) { // 检查 alertInstance 是否存在且有 value 属性
                alertInstance.value.close(); // 关闭提示框
              }
            }, 1 * 60 * 1000); // 1分钟，以毫秒为单位
          } else {
            console.log("执行失败");
          }
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    },
    getData() {
      axios.post('/selectallfile', null)
        .then(response => {
          // 清空选中行数据
          this.selectedRows = [];

          // 将返回的数据设置给 message 变量
          this.message = response.data.message.map(item => {
            if (item.status === true) {
              this.selectedRows.push(item); // 如果 status 为 true，则将行添加到选中数组中
            }
            return item;
          });
          console.log(this.message); // 打印检查数据
        })
        .catch(error => {
          console.error('Error sending data:', error);
        });
    },
    handlePageClick(pageNum) {
      this.currentPage = pageNum; // 处理页码点击事件
    },
    updateMaxprice(row) {
      console.log(row.Wordid, row.Setprice);
      axios.post('/update', { Wordid: row.Wordid, Setprice: row.Setprice })
        .then(response => {
          console.log('最高价格已更新');
        })
        .catch(error => {
          console.error('更新最高价格时出错：', error);
        });
    },
    updateNowprice(row) {
      console.log(row.Wordid, row.NowPrice);
      axios.post('/updatenowprice', { Wordid: row.Wordid, Nowprice: row.NowPrice })
        .then(response => {
          console.log('当前价格已更新');
        })
        .catch(error => {
          console.error('更新当前价格时出错：', error);
        });
    },
    handleSwitchChange(row) {
      const dataToSend = {
        wordid: row.Wordid,
        status: row.Status,
      };
      console.log(dataToSend);
      // 发送到后端的逻辑，可以使用 axios 
      axios.post('/updateStatus', dataToSend, {
        headers: {
          'Content-Type': 'application/json' // 告诉后端发送的是 JSON 格式的数据
        }
      })
        .then(response => {
          console.log('状态更新成功');
        })
        .catch(error => {
          console.error('状态更新失败', error);
        });
    },
    updateUpBySuggest(row) {
      console.log(row.Wordid, row.UpBySuggest);
      axios.post('/updateupbysuggest', { Wordid: row.Wordid, UpBysuggest: row.UpBySuggest })
        .then(response => {
          console.log('追加价格已更新');
        })
        .catch(error => {
          console.error('追加最高价格时出错：', error);
        });
    },


  },

  mounted() {
    setInterval(this.fetchData, 20 * 60 * 1000);
    this.getData();

  },
});
</script>


<style scoped></style>