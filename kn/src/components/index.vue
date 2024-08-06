<template>
    <div class="common-layout">
      <el-container>
        <el-header>Header</el-header>
        <el-container>
          <el-aside width="10%">
            <el-row class="tac">
                <el-menu active-text-color="#ffd04b" background-color="#545c64" class="el-menu-vertical-demo"
                  default-active="0" text-color="#fff" router>
                  <el-menu-item index="table" >
                    <el-icon>
                      <HomeFilled />
                    </el-icon>
                    <span>推广</span>
                  </el-menu-item>
                </el-menu>
             
            </el-row>
          </el-aside>
          <el-main>
            <router-view />
          </el-main>
        </el-container>
      </el-container>
    </div>
  </template>
  
  <script>
  
  </script>
  <style>
  .el-header {
    background-color: #373d41;
  }
  
  .el-aside {
    background-color: #545c64;
    height: 100vh;
  }
  .el-menu-item{
    width: 188px;
  }

  </style>