<template>
    <div>
        {{ totalMem }} {{ freeMem }} {{ availableMem }}
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';

const totalMem = ref(0);
const freeMem = ref(0);
const availableMem = ref(0);

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

</style>