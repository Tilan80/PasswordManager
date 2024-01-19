<script setup>
import { ref } from 'vue';
import { router } from '../router';
import { InsertDocument } from '../../wailsjs/go/main/App';

// Reactive data
const platform = ref('');
const username = ref('');
const isSaved = ref(false);

// Function to save the data
function saveData() {
    // Call the InsertOwn function with the input values
    InsertDocument(platform.value, username.value)
        .then(() => {
            // Data saved successfully
            isSaved.value = true;
            // Clear the input fields
            platform.value = '';
            username.value = '';
        })
        .catch((error) => {
            // Handle error if the data couldn't be saved
            console.error('Error saving data:', error);
            isSaved.value = false;
        });
}
</script>

<template>
    <div>
        <h1>Add new save</h1>

        <!-- Button to go back -->
        <v-btn class="mx-auto bg-green-accent-1 font-weight-bold ma-6" size="x-large" rounded="xl"
            @click="router.push('/Home')">Go back</v-btn>

        <!-- Text fields for platform and username -->
        <v-text-field v-model="platform" label="Platform" outlined></v-text-field>
        <v-text-field v-model="username" label="Username" outlined></v-text-field>

        <!-- Button to save the data -->
        <v-btn class="mx-auto bg-blue font-weight-bold ma-6" size="x-large" rounded="xl"
            @click="saveData">Save</v-btn>

        <!-- Display a message if the document is saved -->
        <v-alert v-if="isSaved" type="success" class="mx-auto" outlined>
            Document has been saved successfully!
        </v-alert>
    </div>
</template>
  

  