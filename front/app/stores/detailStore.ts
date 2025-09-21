import { defineStore } from 'pinia';

export const useDetailStore = defineStore('detail',
    ()=>{
        const detailId = ref(0);
        const detailImage = ref('');
        const detailTime = ref(0);
        const detailDate = ref('');
        const detailComment = ref('');
        
        const setDetail = (id:number,image:string,time:number,date:string,comment:string) => {
            detailId.value = id;
            detailImage.value = image;
            detailTime.value = time;
            detailDate.value = date;
            detailComment.value = comment;
        }
        
        return { detailId,detailImage,detailTime,detailDate,detailComment,setDetail }
        },
);