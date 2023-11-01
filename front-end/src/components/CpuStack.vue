<template>
    <div class="stack-container">
      <div v-for="n in intUsage" class="fixed-size-box" :key="n">
        <div class="cpu-stack"></div>
      </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const usage = defineProps({
    usage: Number
});

const intUsage = computed(() => {
    if (usage && usage.usage !== undefined) {
        const val = Math.round(usage.usage / 10) + 1;
        console.log(val);
        return val;
    }
    // Handle the case where usage.usage is undefined, e.g., return a default value or throw an error.
    // You can replace the following line with your desired behavior.
    return 0; // Return 0 as a default value, for example.
});
</script>

<style scoped>
.stack-container {
  display: flex;
  flex-direction: column-reverse;
  align-items: left;
  height: 400px; /* Set a fixed height for the container */
}

.fixed-size-box {
  width: 50px; /* Set a fixed width for the boxes */
  margin: 5px; /* Add some spacing between the boxes */
}

.cpu-stack {
  width: 100%;
  height: 20px; /* Set a fixed height for the CPU stack divs */
  background-color: #3498db; /* Background color for the CPU stacks */
  border: 1px solid #2980b9; /* Border for the CPU stacks */
  border-radius: 3px; /* Add some border radius for rounded corners */
}
</style>