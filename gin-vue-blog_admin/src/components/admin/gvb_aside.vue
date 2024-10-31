<script setup>

</script>

<template>
    <aside class="gvb_aside">
        <div class="gvb_aside_header flex">
            <div class="gvb_aside_logo">
                <img src="https://www.logoko.com.cn/uploadfile/icon_case/201808/5b86063b48933.png">
            </div>
            <div class="gvb_aside_title">导航</div>
        </div>
        <div class="gvb_aside_body">

        </div>
    </aside>
</template>

<style lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);

  .gvb_aside_header{
    width: 100%;
    height: 260px;
    flex-direction: column;
  }

  .gvb_aside_logo{
    padding: 10px;
    img{
        width: 90%;
    }
  }

  .gvb_aside_title{
    margin-top: 20px;
    font-size: 20px;
  }
}
</style>