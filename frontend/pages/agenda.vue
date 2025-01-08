<template>
    <h1>{{ word }}</h1>
</template>

<script setup lang="ts">
    const word = ref("")
    onMounted(() => {
        const route = useRoute()
        word.value = <string> route.query.word
    })
</script>

<style scoped>
</style>
