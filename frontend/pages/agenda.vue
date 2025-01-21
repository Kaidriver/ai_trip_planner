<template>
    <div class = "container">
        <div class="loading-animation" v-if="state=='loading'">
            <video autoplay>
                <source src="~/assets/loading-animation.webm" type="video/webm">
            </video>
            <h5>Generating...</h5>
        </div>

        <div class="error-message" v-if="state=='error'">
            <h1>Error Loading Agenda</h1>
            <div class = "return-btn">
                <NuxtLink to="/"><button type="button" class="btn error-button">Return Home</button></NuxtLink>
            </div>
        </div>

        <h1 id="header-text"></h1>
        <h1 id="sub-header-text"></h1>

        <div v-if="state=='loaded'" class = "row content-block">
            <div class="days-display">
                <h5 v-for="(item, index) in agenda" @click="currentDay = agenda[index]">Day {{ index + 1 }}</h5>
            </div>
            <div class = "col-md-3 meals-col agenda-col">
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
            <div class = "col-md-6 activities-col agenda-col">
                <div class = "row">
                    <h5 class = "activities-header">Morning</h5>
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
                    <h5 class = "activities-header">Afternoon</h5>
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
                    <h5 class = "activities-header">Night</h5>
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
            <div class = "col-md-3 agenda-col">
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
</template>

<script setup lang="ts">
    import { DotLottieVue } from '@lottiefiles/dotlottie-vue'
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
    const state = ref("")
    const route = useRoute()

    function timeout(ms: number) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    onMounted(async () => {
        if (Object.hasOwn(route.query, 'destination') && Object.hasOwn(route.query, 'budget') &&
            Object.hasOwn(route.query, 'activityLevel') && Object.hasOwn(route.query, 'radiusMiles') && 
            Object.hasOwn(route.query, 'numDays')) {
            state.value = "loading"
            console.log("url: " + config.public.prod_endpoint)
            const { data, status, error, refresh, clear } = await useFetch(config.public.prod_endpoint, {
                method: "GET",
                query: route.query,
                headers: {"x-api-key": config.public.secret},
            })

            if (error.value != undefined) {
                state.value = "error"
            }
            else {
                agenda.value = <DayAgenda[]> data.value
                currentDay.value = agenda.value[0]

                state.value = "startLoading"
                
                let headerText = 'Welcome to ' + route.query.destination;
                
                let headerTextElements = headerText.split("").map((c) => {
                    let spanElement = document.createElement('span')
                    spanElement.appendChild(document.createTextNode(c))
                    spanElement.style.display = 'hidden'
                    return spanElement
                });
                
                let el = document.getElementById("header-text");
        
                await (async () => {
                    for (let i = 0; i < headerTextElements.length; i++) {
                        let e = headerTextElements[i]
                        el!.append(e);
                        e.style.display = 'initial'
                        
                        await timeout(85)
                    }
                })();

                let subElText = 'Here is your ' + agenda.value.length + ' day itinerary';
                console.log(agenda)

                let subElTextElements = subElText.split("").map((c) => {
                    let spanElement = document.createElement('span')
                    spanElement.appendChild(document.createTextNode(c))
                    spanElement.style.display = 'hidden'
                    return spanElement
                });

                let subEl = document.getElementById("sub-header-text");

                await (async () => {
                    for (let i = 0; i < subElTextElements.length; i++) {
                        let e = subElTextElements[i]
                        subEl!.append(e);
                        e.style.display = 'initial'
                        
                        await timeout(85)
                    }
                })();
                
                state.value = "loaded"
            }
        }
        else {
            state.value = 'error'
        }
    })
    

    watch(currentDay, () => {
        document.querySelectorAll('.agenda-col').forEach((col) => {
            (<HTMLElement> col).classList.remove('col-transition');
            void (<HTMLElement> col).offsetWidth;
            (<HTMLElement> col).classList.add('col-transition')
        })
    })
</script>
!

<style scoped>

h1 {
    font-family: 'Readex Pro';
    font-size: 500%;
    background-image: linear-gradient(270deg, #FF417A 40%, #5596FF 85%);
    background-clip: text;
    color: transparent;
}

.loading-animation {
    margin: 0;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 12em;
}

.loading-animation h5 {
    position: relative;
    left: 10%;
    white-space: nowrap; /* Keeps the content on a single line */
    background-image: linear-gradient(270deg, #FF417A 40%, #5596FF 85%);
    background-clip: text;
    color: transparent;
    font-family: 'Readex Pro';
    animation: typing 2.5s steps(40, end) infinite
}

@keyframes typing {
  from { width: 0 }
  to { width: 100% }
}

@keyframes fadeInOpacity {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}

.show-text {
    display: initial;
}

.card-body, .activities-col {
    font-family: 'Mate';
}

.card {
  position: relative;
  backdrop-filter: blur(10px);
  border: none;
}

.card::before {
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

.btn {
    width: 100%;
}

.return-btn {
  position: relative;
  backdrop-filter: blur(10px);
  padding: 4px;
  width: 10em;
  transition: 0.2s;
  margin: 0 auto;
  margin-top: 3%;
}

.return-btn::before {
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

.return-btn:hover {
    transform: scale(1.1);
}

.error-message {
    margin: 0;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
    width: 100%;
}

.error-message h1 {
    font-size: 500%;
}

.error-button {
    text-decoration: none;
    background-image: linear-gradient(270deg, #FF417A 40%, #5596FF 85%);
    background-clip: text;
    color: transparent;
}

.meals-col .card {
    margin-bottom: 10%;
}

.activities-col .row {
    margin-bottom: 4%;
}

.activities-col .card {
    height: 100%;
}

.activities-header {
    font-size: 150%;
    font-weight: 500;
}

.days-display {
    display: flex;  
    gap: 1em;
}

.days-display h5 {
    font-size: 1em;
    font-family: 'Readex Pro';
    background-image: linear-gradient(270deg, #FF417A 40%, #5596FF 85%);
    background-clip: text;
    color: transparent;
    transition: 0.2s;
}

.days-display h5:hover {
    cursor: pointer;
    transform: scale(1.1);
    opacity: 0.8;
}

.content-block {
    margin-top: 2%;

    animation-name: fadeInOpacity;
	animation-iteration-count: 1;
	animation-timing-function: ease-in;
	animation-duration: 0.5s;
}

.col-transition {
    animation-name: fadeInOpacity;
	animation-iteration-count: 1;
	animation-timing-function: ease-in;
	animation-duration: 0.5s;
}

.activities-col .row .col-md-6:not(:nth-child(2), :nth-child(3)) {
    margin-top: 3%;
}

@media screen and (max-width: 1000px) {
    h1 {
        font-size: 400%;
    }
}

@media screen and (max-width: 768px) {
    h1 {
        font-size: 350%;
    }

    .activities-col .row .col-md-6:not(:first-child) {
        margin-top: 3%;
    }
}
</style> 
