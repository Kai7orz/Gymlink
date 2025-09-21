// Import the functions you need from the SDKs you need
import { initializeApp, getApps, getApp } from "firebase/app";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
export default defineNuxtPlugin(() => {
const firebaseConfig = {
  apiKey: "AIzaSyDLENAgDartCqlupd9YZTVAylQ6XQ8Ivw4",
  authDomain: "rizap-hackathon.firebaseapp.com",
  projectId: "rizap-hackathon",
  storageBucket: "rizap-hackathon.firebasestorage.app",
  messagingSenderId: "912445110020",
  appId: "1:912445110020:web:29730c6972015621508251",
  measurementId: "G-KEET8WYBC8"
};

const app = !getApps().length ? initializeApp(firebaseConfig) : getApp();
const auth = getAuth(app);

return {
  provide: {
    firebaseAuth: auth,
  },
};
});