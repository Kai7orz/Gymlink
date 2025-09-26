import { defineStore } from 'pinia';

export const useDetailStore = defineStore('detail',
    ()=>{
        const detailId = ref(0);
        const detailName = ref('')
        const detailImage = ref('');
        const detailTime = ref(0);
        const detailDate = ref('');
        const detailComment = ref('');
        const detailLikesCount = ref(0);
        
        const setDetail = (id:number,name:string,image:string,time:number,date:string,comment:string,likesCount:number) => {
            detailId.value = id;
            detailName.value = name;
            detailImage.value = image;
            detailTime.value = time;
            detailDate.value = date;
            detailComment.value = comment;
            detailLikesCount.value = likesCount;
        }
        
        return { detailId,detailImage,detailTime,detailDate,detailComment,detailLikesCount,setDetail }
        },
);