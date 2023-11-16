<template>
  <div class="stack-container">
    <div v-for="n in intUsage" class="fixed-size-box" :key="n">
      <div class="cpu-stack" :id="'box' + n"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed} from 'vue';

const usage = defineProps({
  usage: Number
});

const intUsage = computed(() => {
  if (usage && usage.usage !== undefined) {
    const val = Math.floor(usage.usage / 10) + 1;
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
  border-radius: 3px; /* Add some border radius for rounded corners */
  background-color: #d04708;
  border-color: #d04708;
}

#box1 {
  background-color: white;
  border-color: white;
}

#box2 {
  background-color: #ffcccc; /* Very light red */
  border-color: #ffcccc;
}

#box3 {
  background-color: #ff9999; /* Light red */
  border-color: #ff9999;
}

#box4 {
  background-color: #ff6666; /* Red */
  border-color: #ff6666;
}

#box5 {
  background-color: #ff3333; /* Slightly darker red */
  border-color: #ff3333;
}

#box6 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

#box7 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

#box8 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

#box9 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

#box10 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

#box11 {
  background-color: #ff0000; /* Strong red */
  border-color: #ff0000;
}

</style>