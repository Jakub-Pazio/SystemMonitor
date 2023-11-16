<template>
  <div class="outer">
    <h2>Temp Check</h2>
    <p :class=" { hot: isHot }">CPU temperature: {{ temp / 1000 }}</p>
  </div>
</template>

<script setup lang="ts">
import {onBeforeUnmount, onMounted, ref} from 'vue';

const temp = ref(0);
const isHot = ref(false);

async function getTemp() {
  try {
    const res = await fetch('/temp');
    const data = await res.json();
    console.log(data);
    if (data && data.name === 'temp' && typeof data.value === 'number') {
      if (data.value > 50000) {
        isHot.value = true;
      } else {
        isHot.value = false;
      }
      temp.value = (data.value); // Update usage with the value as a percentage
    } else {
      console.error('Invalid data format received from the API.');
      temp.value = 0; // Return a default value
    }
  } catch (error) {
    console.error('Error fetching CPU usage:', error);
    temp.value = 0; // Return a default value in case of an error
  }
}

onMounted(() => {
  // Call getCpuUsage every second
  const updateInterval = setInterval(getTemp, 1000);

  // Stop the interval when the component is unmounted (cleanup)
  onBeforeUnmount(() => {
    clearInterval(updateInterval);
  });
});
</script>

<style scoped>
.hot {
  color: red;
}

.outer {
  padding-top: 10px;
}
</style>