<template>
    <h1>Agenda:</h1>

    <p>{{ agendaText }}</p>
</template>

<script setup lang="ts">
    const config = useRuntimeConfig()
    const agendaText = ref("Loading...")

    onMounted(() => {
       const route = useRoute()
       
       useFetch(config.public.prod_endpoint, {
            method: "GET",
            query: route.query,
            headers: {"x-api-key": config.public.secret},
        }).then((data) => {
            agendaText.value = <string> data.data.value
        }).catch((error) => {
            agendaText.value = "Error generating schedule"
        })
    })
</script>

<style scoped>
</style>
