<template>
    <div class = "container">
        <h1>Plan My Trip AI</h1>

        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
            <UFormGroup label="Destination" name="destination">
                <UInput v-model="state.destination" />
            </UFormGroup>

            <UFormGroup label="Budget" name="budget">
                <USelect v-model="state.budget" :options="levels" placeholder=""/>
            </UFormGroup>

            <UFormGroup label="Activity Level" name="activityLevel">
                <USelect v-model="state.activityLevel" :options="levels" placeholder=""/>
            </UFormGroup>

            <UFormGroup label="Activity Type" name="activityTypes">
                <USelectMenu v-model="state.activityTypes" :options="activities" multiple placeholder="Select Preferences" />
            </UFormGroup>

            <UFormGroup label="Resturant Type" name="activityTypes">
                <USelectMenu v-model="state.resturantTypes" :options="resturants" multiple placeholder="Select Preferences" />
            </UFormGroup>

            <UFormGroup label="Radius" name="radiusMiles">
                <UInput v-model="state.radiusMiles" type="number" placeholder="Maximum 500" />
            </UFormGroup>

            <UFormGroup label="Number Days" name="numDays">
                <UInput v-model="state.numDays" type="number"  placeholder="Maximum 7"/>
            </UFormGroup>

            <UButton type="submit">
                Submit
            </UButton>
        </UForm>
    </div>
</template>

<script setup lang="ts">
    import { object, number, string, array, type InferType } from 'yup'
    import type { FormSubmitEvent } from '#ui/types'

    const schema = object({
        radiusMiles: number().min(0).max(500).required(),
        numDays: number().min(1).max(7).required(),
        destination: string().required(),
        budget: string().required(),
        activityLevel: string().required(),
        activityTypes: array(),
        resturantTypes: array(),
    })

    type Schema = InferType<typeof schema>
    
    const levels = ['Low', 'Medium', 'High']
    const activities = ['Shopping', 'Outdoors', 'Nightlife', 'Museums', 'Thrill']
    const resturants = ['Chinese', 'Mexican', 'American', 'Italian', 'Korean', 'Japanese', 'Indian', 'Thai']

    const state = reactive({
        destination: undefined,
        budget: undefined,
        activityLevel: undefined,
        activityTypes: undefined,
        resturantTypes: undefined,
        radiusMiles: undefined,
        numDays: undefined,
    })

    async function onSubmit(event: FormSubmitEvent<Schema>) {
        // Do something with event.data
        navigateTo({
            path: '/agenda',
            query: event.data
        })
    }
</script>

<style scoped>
</style>
