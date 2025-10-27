<script setup lang="ts">
import type { RecordType } from "~/type";

const props = defineProps<{
  recordList: RecordType[];
  isOwner: boolean;
}>();

const emits = defineEmits<{
  detail: [id: number];
  like: [id: number];
}>();

const toDetail = (id: number) => {
  emits("detail", id);
};

const toDelete = (id: number) => {
  emits("delete", id);
};

const toAccount = (uid: number) => {
  emits("account", uid);
};

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
          :id="record.id"
          class="w-100"
          :is-owner="props.isOwner"
          :user-id="Number(record.user_id)"
          :user-name="record.user_name"
          :image="record.presigned_image"
          :time="record.clean_up_time"
          :date="record.clean_up_date"
          :comment="record.comment"
          :likes-count="record.likes_count"
          @detail="toDetail"
          @delete="toDelete"
          @account="toAccount"
        />
      </v-col>
    </v-row>
  </v-container>
</template>
