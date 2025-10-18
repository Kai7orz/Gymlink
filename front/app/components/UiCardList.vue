<script setup lang="ts">
import type { RecordType } from '~/type';

    const props = defineProps<{
        recordList: RecordType[]
    }>();

    const emits = defineEmits<{
        detail: [id: number],
        like: [id: number],
    }>();

    const toDetail = (id:number) => {
        emits('detail',id)
    }

    const toAccount = (uid:number) => {
        emits('account',uid)
    }

</script>
<template>
  <v-container>
    <v-row>
      <v-col
        v-for="(record, index) in props.recordList"
        :key="index"
        cols="12"
        md="6"      
      >
        <ui-card
          class="w-100"
          :id="record.id"
          :userId="Number(record.user_id)"
          :userName="record.user_name"
          :image="record.presigned_image"
          :time="record.clean_up_time"
          :date="record.clean_up_date"
          :comment="record.comment"
          :likesCount="record.likes_count"
          @detail="toDetail"
          @account="toAccount"
        />
      </v-col>
    </v-row>
  </v-container>
</template>
