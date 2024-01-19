<script setup>
import { ref } from 'vue';
import { router } from '../router';
import { InsertOwn } from '../../wailsjs/go/main/App';

// Reactive data
const platform = ref('');
const username = ref('');
const password = ref('');
const reenteredPassword = ref('');
const isSaved = ref(false);

// Function to save the data
function saveData() {
    // Check if passwords match
    if (password.value === reenteredPassword.value) {
        // Call the InsertOwn function with the input values
        InsertOwn(platform.value, username.value, password.value)
            .then(() => {
                // Data saved successfully
                isSaved.value = true;
                // Clear the input fields
                platform.value = '';
                username.value = '';
                password.value = '';
                reenteredPassword.value = '';
            })
            .catch((error) => {
                // Handle error if the data couldn't be saved
                console.error('Error saving data:', error);
            });
    } else {
        // Show an error message if passwords don't match
        console.error('Passwords do not match!');
    }
}
</script>

<template>
    <div>
        <h1>Add new save with own password</h1>

        <!-- Button to go back -->
        <v-btn class="mx-auto bg-green-accent-1 font-weight-bold ma-6" size="x-large" rounded="xl"
            @click="router.push('/Home')">Go back</v-btn>

        <!-- Text fields for platform, username, and password -->
        <v-text-field v-model="platform" label="Platform" outlined></v-text-field>
        <v-text-field v-model="username" label="Username" outlined></v-text-field>
        <v-text-field v-model="password" label="Password" outlined type="password"></v-text-field>

        <!-- Text field to re-enter password -->
        <v-text-field v-model="reenteredPassword" label="Re-enter Password" outlined type="password"></v-text-field>

        <!-- Button to save the data -->
        <v-btn class="mx-auto bg-blue font-weight-bold ma-6" size="x-large" rounded="xl"
            @click="saveData">Save</v-btn>

        <!-- Display a message if the document is saved -->
        <v-alert v-if="isSaved" type="success" class="mx-auto" outlined>
            Document has been saved successfully!
        </v-alert>
    </div>
</template>
  

  