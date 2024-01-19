<script setup>
import { ref } from 'vue';
import { router } from '../router';
import { DeleteByPlatform } from '../../wailsjs/go/main/App';

// Reactive data
const platform = ref('');
const isDeletionSuccessful = ref(null);

// Function to delete the password by platform
function deleteByPlatform() {
  // Call the DeleteByPlatform function with the input platform
  DeleteByPlatform(platform.value).then(success => {
    // Set the result of the deletion
    isDeletionSuccessful.value = success;

    // You can also navigate back to the Home page after deletion
    router.push('/Home');
  });
}
</script>

<template>
  <div>
    

    <h1>Delete the saved password by platform</h1>

    <!-- Button to go back -->
    <v-btn class="mx-auto bg-green-accent-1 font-weight-bold ma-6" size="x-large" rounded="xl"
      @click="router.push('/Home')">Go back</v-btn>

    <!-- Text field for inputting the platform -->
    <v-text-field v-model="platform" label="Input the platform you want to delete" outlined></v-text-field>

    <!-- Button to delete the password by platform -->
    <v-btn class="mx-auto bg-red accent-1 font-weight-bold ma-6" size="x-large" rounded="xl" 
    @click="deleteByPlatform">Delete By Platform</v-btn>

    <!-- Display a success or error message using v-alert -->
    <v-alert v-if="isDeletionSuccessful === true" type="success" class="mx-auto" outlined>
      Deletion was successful!
    </v-alert>
    <v-alert v-if="isDeletionSuccessful === false" type="error" class="mx-auto" outlined>
      Deletion failed. Please try again.
    </v-alert>
  </div>
</template>
