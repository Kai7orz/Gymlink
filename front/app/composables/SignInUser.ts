import { signInWithEmailAndPassword, type Auth } from "firebase/auth";
import { FirebaseError } from '@firebase/util'

export const signIn = async (email:string,password:string)=>{
    const { $firebaseAuth } = useNuxtApp();
    const auth = $firebaseAuth as Auth;

    try{
        await signInWithEmailAndPassword(auth,email,password)
    } catch(e) {
        if ( e instanceof FirebaseError ) {
            console.error('Firebase Login Error', e.code,e.message);
            throw e
        }
    }
}