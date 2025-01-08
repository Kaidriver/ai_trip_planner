<template>
    <div class = "container">
        <div v-if="Object.keys(currentDay).length !== 0">
            <h1>Welcome to Las Vegas. Here is your 3 day itinerary</h1>
            <div class="days-display">
                <h5 v-for="(item, index) in agenda" @click="currentDay = agenda[index]">Day {{ index + 1 }}</h5>
            </div>

            <div class = "row">
                <div class = "col-md-3 meals-col">
                    <div class = "card">
                        <div class="card-body">
                            <h5 class="card-title">Breakfast</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{{currentDay.breakfast.name}}</h6>
                            <p class="card-text">{{currentDay.breakfast.description}}</p>
                        </div>
                    </div>
                    <div class = "card">
                        <div class="card-body">
                            <h5 class="card-title">Lunch</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{{currentDay.lunch.name}}</h6>
                            <p class="card-text">{{ currentDay.lunch.description }}</p>
                        </div>
                    </div>
                    <div class = "card">
                        <div class="card-body">
                            <h5 class="card-title">Dinner</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{{currentDay.dinner.name}}</h6>
                            <p class="card-text">{{ currentDay.dinner.description }}</p>
                        </div>
                    </div>
                </div>
                <div class = "col-md-6 activities-col">
                    <div class = "row">
                        <h5>Morning</h5>
                        <div class="col-md-6" v-for="activity in currentDay.morningActivities">
                            <div class="card">
                                <div class="card-body">
                                    <h5 class="card-title">{{activity.name}}</h5>
                                    <p class="card-text">{{activity.description}}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class = "row">
                        <h5>Afternoon</h5>
                        <div class="col-md-6"  v-for="activity in currentDay.afternoonActivities">
                            <div class="card">
                                <div class="card-body">
                                    <h5 class="card-title">{{activity.name}}</h5>
                                    <p class="card-text">{{activity.description}}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class = "row">
                        <h5>Night</h5>
                        <div class="col-md-6" v-for="activity in currentDay.nightActivities">
                            <div class="card">
                                <div class="card-body">
                                    <h5 class="card-title">{{activity.name}}</h5>
                                    <p class="card-text">{{activity.description}}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class = "col-md-3">
                    <div class = "card">
                        <div class="card-body">
                            <h5 class="card-title">Hotel</h5>
                            <h6 class="card-subtitle mb-2 text-muted">{{ currentDay.hotel.name }}</h6>
                            <p class="card-text">{{currentDay.hotel.description}}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    interface Activity {
        name: string,
        description: string,
    }

    interface DayAgenda {
        day: number,
        hotel: Activity,
        breakfast: Activity,
        morningActivities: Activity[],
        lunch: Activity,
        afternoonActivities: Activity[],
        dinner: Activity,
        nightActivities: Activity[],
    }

    const config = useRuntimeConfig()
    const agenda = ref<DayAgenda[]>([])

    const currentDay = ref<DayAgenda>(<DayAgenda> {})
    
    onMounted(() => {
       const route = useRoute()
       
       useFetch(config.public.prod_endpoint, {
            method: "GET",
            query: route.query,
            headers: {"x-api-key": config.public.secret},
        }).then((data) => {
            agenda.value = <DayAgenda[]> data.data.value
            currentDay.value = agenda.value[0]
        }).catch((error) => {
            agenda.value = []
        })
    })
</script>
!

<style scoped>
.meals-col .card {
    margin-bottom: 10%;
}

.activities-col .row {
    margin-bottom: 8%;
}

.days-display {
    display: flex;  
    gap: 1em;
}

.days-display h5 {
    font-size: 1em;
}

.days-display h5:hover {
    cursor: pointer;
    text-decoration: underline;
    opacity: 0.8;
}
</style> 
