<script setup lang="ts">
    const emits = defineEmits<{
        close: [],
    }>();

    const isShownMenu = defineModel('isShownMenu')
</script>

<template>
    <v-dialog v-model="isShownMenu" max-width="500">
        <v-card>
            <v-card-title class="text-h5">運動記録</v-card-title>
            <v-card-text>
                
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="emits('close')">Close</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
