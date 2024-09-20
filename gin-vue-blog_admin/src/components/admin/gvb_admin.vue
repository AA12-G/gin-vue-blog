<script setup>
import { ref } from 'vue';
import {useStore} from "@/stores/store"

const store=useStore()
store.loadTheme()

function setTheme(){
   store.setTheme()
}
</script>

<template>
    <i v-if="store.theme" class="fa fa-sun-o" @click="setTheme"></i>
    <i v-else class="fa fa-moon-o" @click="setTheme"></i>
</template>

<style>

</style>