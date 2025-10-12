<script setup lang="ts">
import type { ExerciseRecordType } from '~/type';

    const props = defineProps<{
        exerciseMocksList: ExerciseRecordType[]
    }>();

    const emits = defineEmits<{
        detail: [id: number],
        like: [id: number],
    }>();

    const toDetail = (id:number) => {
        emits('detail',id)
    }

    const like = (id:number) => {
        emits('like',id)
    }

    const toAccount = (uid:number) => {
        emits('account',uid)
    }
</script>
<template>
  <v-container>
    <v-row>
      <v-col
        v-for="(exerciseMock, index) in props.exerciseMocksList"
        :key="index"
        cols="12"
        md="6"      
      >
        <ui-card
          class="w-100"
          :id="exerciseMock.id"
          :userId="Number(exerciseMock.user_id)"
          :userName="exerciseMock.user_name"
          :image="exerciseMock.exercise_image"
          :time="exerciseMock.exercise_time"
          :date="exerciseMock.exercise_date"
          :comment="exerciseMock.comment"
          :likesCount="exerciseMock.likes_count"
          @detail="toDetail"
          @like="like"
          @account="toAccount"
        />
      </v-col>
    </v-row>
  </v-container>
</template>
