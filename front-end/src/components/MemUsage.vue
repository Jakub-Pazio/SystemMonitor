<template>
    <div class="outer">
        <Pie :data="chartData" :options="chartOptions"/>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, computed } from 'vue';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Pie } from 'vue-chartjs'

const totalMem = ref(0);
const freeMem = ref(0);
const availableMem = ref(0);

const chartData = computed(() => ({
    labels: ['Used', 'Free', 'Available'],
    datasets: [
        {
            label: 'Memory Usage',
            backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
            data: [totalMem.value - availableMem.value, freeMem.value, availableMem.value]
        }
    ]
}));

const chartOptions = ref({
    responsive: true,  // Set to true to allow the chart to resize
    maintainAspectRatio: false, // Set to false to customize aspect ratio
    width: 200, // Set the width of the chart
    height: 200, // Set the height of the chart
});

ChartJS.register(ArcElement, Tooltip, Legend)

async function getMemUsage() {
    try {
        const res = await fetch('/mem');
        const data = await res.json();

        console.log(data);
        totalMem.value = data.total
        freeMem.value = data.free
        availableMem.value = data.avail
    } catch (error) {
        console.error('Error fetching CPU usage:', error);
        // usage.value = 0;
    }
}

onMounted(() => {
    // Call getCpuUsage every second
    const updateInterval = setInterval(getMemUsage, 1000);

    // Stop the interval when the component is unmounted (cleanup)
    onBeforeUnmount(() => {
        clearInterval(updateInterval);
    });
});

</script>

<style scoped>
.outer{
    padding: 20px;
}
</style>