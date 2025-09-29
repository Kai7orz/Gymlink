import { createUserWithEmailAndPassword, getAuth, type Auth } from "firebase/auth";
import { FirebaseError } from '@firebase/util'

export const signUp = async (email:string,password:string)=>{
    const { $firebaseAuth } = useNuxtApp();
    const auth = $firebaseAuth as Auth;

    try{
        const userCredential = await createUserWithEmailAndPassword(auth,email,password)
        console.log("successful サインアップ：",userCredential.user);
    } catch(e) {
        if ( e instanceof FirebaseError ) {
            console.error('FirebaseError', e.code,e.message);
        }
        throw e
    }
}