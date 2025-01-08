<template>
    <h1>hello</h1>
    <button @click="navigate" class="btn btn-primary"> Page  </button>
</template>

<script setup lang="ts">
    const navigate = (event) => {
        alert("HI")
        navigateTo({
            path: '/agenda',
            query: {
                word: "pineapple"
            }
        })
    }
</script>

<style scoped>
</style>
