import { createUserWithEmailAndPassword, type Auth } from "firebase/auth";
import { FirebaseError } from "@firebase/util";

export const signUp = async (email: string, password: string) => {
  const { $firebaseAuth } = useNuxtApp();
  const auth = $firebaseAuth as Auth;

  try {
    await createUserWithEmailAndPassword(auth, email, password);
  }
  catch (e) {
    if (e instanceof FirebaseError) {
      console.error("FirebaseError", e.code, e.message);
    }
    throw e;
  }
};
