import { signOut } from 'firebase/auth';

export const signOutUser = async() => {
    try {
        const { $firebaseAuth } = useNuxtApp();
        await signOut($firebaseAuth)
    } catch(e) {
        console.log("signout error ",e)
    }
}