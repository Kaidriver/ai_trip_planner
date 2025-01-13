<template>
    <div class = "container">
        <div class = "header">
            <h1>Triply.AI</h1>
            <h3>Auto-generate your next trip ✈️</h3>
        </div>
        
        <div class = "input-body">
                <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit" @error="onError">
                <div class = "input-row">
                    <span>Let's go to</span>
                    <div class = "gradient-input input-location">
                        <UInput v-model="state.destination" color="primary" variant="none" class="input-item"/>
                    </div>  

                    <div class="break"></div>
                    <span>for</span>
                    <div class = "gradient-input input-days">
                        <UInput v-model="state.numDays" color="primary" variant="none" type="number" class="input-item"/>
                    </div>
                    <span>days, within</span>
                    <div class = "gradient-input input-miles">
                        <UInput v-model="state.radiusMiles" color="primary" variant="none" type="number" class="input-item" />
                    </div>
                    <span>miles</span>
                </div>
                
                <div class = "input-row">
                    <span>I prefer to do</span>
                    <div class = "gradient-input input-activity-types">
                        <USelectMenu v-model="state.activityTypes" :options="activities" variant="none" multiple/>
                    </div>
                    <span v-if="width > 800">activities and go to</span>
                    <span v-if="width <= 800">activities</span>
                    <div class="break"></div>
                    <span v-if="width <= 800">and go to</span>
                    <div class = "gradient-input input-resturant">
                        <USelectMenu v-model="state.resturantTypes" :options="resturants" variant="none" multiple/>
                    </div>
                    <span>resturants</span>
                </div>

                <div class = "input-row">
                    <span>I want a</span>
                    <div class = "gradient-input input-activity-level">
                        <USelect v-model="state.activityLevel" :options="levels" variant="none" placeholder=" "/>
                    </div>
                    <span v-if="width > 800">agenda with a</span>
                    <span v-if="width <= 800">agenda</span>
                    <div class="break"></div>
                    <span v-if="width <= 800">with a</span>
                    <div class = "gradient-input input-budget">
                        <USelect v-model="state.budget" :options="levels" variant="none" placeholder=" "/>
                    </div>
                    <span>budget</span>
                </div>

                <div class = "gradient-input input-submit-btn">
                    <UButton type="submit" color="white" variant="ghost" class="submit-btn" block>
                        Submit
                    </UButton>
                </div>

                <h5>{{ errorText }}</h5>
            </UForm>
        </div>
    </div>
</template>

<script setup lang="ts">
    import { object, number, string, array, type InferType } from 'yup'
    import type { FormSubmitEvent } from '#ui/types'
    import { useWindowSize } from '@vueuse/core'

    const { width } = useWindowSize()

    const schema = object({
        radiusMiles: number().min(0).max(500).required(),
        numDays: number().integer().min(1).max(7).required(),
        destination: string().max(50).required(),
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
        activityLevel: "",
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

    const errorText = ref("")
    async function onError(event: ErrorEvent) {
        console.log(event)
        errorText.value = "Destination length must be less than 50 characters, days less than 8 and miles less than 500"
    }
</script>

<style scoped>
h1 {
    font-family: 'Readex Pro';
}

h3, h5 {
    font-family: 'Mate';
}

/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none !important;
  margin: 0;
}
    
/* Firefox */
input[type=number] {
  -moz-appearance: textfield !important;
}

.header h1 {
    font-size: 800%;
    background-image: linear-gradient(270deg, #FF417A 40%, #5596FF 70%);
    background-clip: text;
    color: transparent;
}

.header {
    text-align: center;
    margin-top: 5%;
}

.input-body {
    margin-top: 6%;
}

.input-body span {
    font-size: 150%;
}

.input-body h5 {
    text-align: center;
    font-size: 100%;
    color: red;
}

.input-row {
    display: flex;
    justify-content: center;
    gap: 1.5%;
    flex-wrap: wrap;
}

.input-item {
    font-size: 20%;
}

.input-days {
    width: 5%;
}

.input-miles {
    width: 8%;
}

.input-activity-level, .input-budget {
    width: 10%;
}

.input-resturant, .input-activity-types{
    width: 12%;
    z-index: 1;
}

.input-submit-btn {
    margin: 0 auto;
    width: 10%;
    transition: 0.2s;
}

.input-submit-btn:hover {
    transform: scale(1.1);
}

.submit-btn {
    background-image: linear-gradient(90deg, #FF417A 40%, #5596FF 70%);
    background-clip: text;
    color: transparent;
}

.gradient-input {
  position: relative;
  backdrop-filter: blur(10px);
  padding: 4px;
}

.gradient-input::before {
  content: "";
  position: absolute;
  inset: 0;
  border: 2px solid transparent;
  border-radius: 12px;
  background: linear-gradient(140deg, #FF417A 40%, #5596FF 70%) border-box;
  mask: linear-gradient(#fff 0 0) padding-box, linear-gradient(#fff 0 0);
  mask-composite: exclude;
  z-index: -1;
}

.break {
    display: none;
}

@media screen and (max-width: 1200px) {
    .input-days {
        width: 6%;
    }

    .input-miles {
        width: 8%;
    }

    .input-activity-level, .input-budget {
        width: 13%;
    }

    .input-resturant, .input-activity-types{
        width: 13%;
        z-index: 1;
    }
}

@media screen and (max-width: 800px) {
    .container {
        width: 100% !important;
        max-width: 100% !important;
    }

    .input-body span {
        font-size: 125%;
    }

    .break {
        display: initial;
        flex-basis: 100%;
        height: 0;
        margin-top: 1.5%;
    }

    .input-days {
        width: 15%;
    }

    .input-miles {
        width: 16%;
    }

    .input-activity-level, .input-budget {
        width: 40%;
    }

    .input-resturant, .input-activity-types{
        width: 35%;
        z-index: 1;
    }

    .input-submit-btn {
        width: 20%;
    }
}

@media screen and (max-width: 600px) {
    .header h1 {
        font-size: 500%;
    }

    .header h3 {
        font-size: 125%;
    }

    .input-body {
        margin-top: 10%;
    }

    .input-submit-btn  {
        width: 30%;
    }
}
</style>
