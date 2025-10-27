import { defineStore } from "pinia";

export const useDetailStore = defineStore("detail",
  () => {
    const detailId = ref(0);
    const detailUserId = ref(0);
    const detailName = ref("");
    const detailPresignedImage = ref("");
    const detailTime = ref(0);
    const detailDate = ref("");
    const detailComment = ref("");
    const detailLikesCount = ref(0);

    const setDetail = (id: number, user_id: number, name: string, presigned_image: string, clean_up_time: number, clean_up_date: string, comment: string, likesCount: number) => {
      detailId.value = id;
      detailUserId.value = user_id;
      detailName.value = name;
      detailPresignedImage.value = presigned_image;
      detailTime.value = clean_up_time;
      detailDate.value = clean_up_date;
      detailComment.value = comment;
      detailLikesCount.value = likesCount;
    };

    return { detailId, detailUserId, detailName, detailPresignedImage, detailTime, detailDate, detailComment, detailLikesCount, setDetail };
  },
);
