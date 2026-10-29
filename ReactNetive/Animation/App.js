import React, { useEffect } from "react";
import { Button, Alert, Platform, View } from "react-native";
import * as Notifications from "expo-notifications";
import * as Device from "expo-device";

// Configure notification behavior
Notifications.setNotificationHandler({
  handleNotification: async () => ({
    shouldShowAlert: true,
    shouldPlaySound: true,
    shouldSetBadge: false,
  }),
});

export default function App() {
  useEffect(() => {
    // Register for push notifications
    registerForPushNotificationsAsync();

    // Listener for notifications received while the app is in the foreground
    const notificationListener = Notifications.addNotificationReceivedListener(
      (notification) => {
        Alert.alert(
          "Notification Received",
          `Title: ${notification.request.content.title}\nBody: ${notification.request.content.body}`
        );
      }
    );

    // Listener for user interaction with a notification
    const responseListener =
      Notifications.addNotificationResponseReceivedListener((response) => {
        Alert.alert(
          "Notification Tapped",
          `Data: ${JSON.stringify(response.notification.request.content.data)}`
        );
      });

    // Cleanup listeners on unmount
    return () => {
      Notifications.removeNotificationSubscription(notificationListener);
      Notifications.removeNotificationSubscription(responseListener);
    };
  }, []);

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
      <Button
        title="Test Push Notification"
        onPress={async () => {
          await sendTestNotification();
        }}
      />
    </View>
  );
}

// Function to register for push notifications
async function registerForPushNotificationsAsync() {
  if (!Device.isDevice) {
    Alert.alert("Must use physical device for Push Notifications");
    return;
  }

  try {
    const { status: existingStatus } =
      await Notifications.getPermissionsAsync();
    let finalStatus = existingStatus;

    if (existingStatus !== "granted") {
      const { status } = await Notifications.requestPermissionsAsync();
      finalStatus = status;
    }

    if (finalStatus !== "granted") {
      Alert.alert("Failed to get push token for push notifications!");
      return;
    }

    const token = (await Notifications.getExpoPushTokenAsync()).data;
    console.log("Push Notification Token:", token);
    Alert.alert("Push Token", token);

    // Configure notification channel for Android
    if (Platform.OS === "android") {
      await Notifications.setNotificationChannelAsync("default", {
        name: "default",
        importance: Notifications.AndroidImportance.MAX,
        vibrationPattern: [0, 250, 250, 250],
        lightColor: "#FF231F7C",
      });
    }

    return token;
  } catch (error) {
    console.error("Error registering for push notifications:", error);
    Alert.alert(
      "Error",
      "An error occurred while setting up push notifications."
    );
  }
}

// Function to send a test notification
async function sendTestNotification() {
  try {
    await Notifications.scheduleNotificationAsync({
      content: {
        title: "📬 Test Notification!",
        body: "This is a test notification from your app.",
        data: { testData: "Some data to test" },
        sound: true,
      },
      trigger: { seconds: 2 }, // Notification delay
    });
  } catch (error) {
    console.error("Error sending test notification:", error);
    Alert.alert("Error", "Failed to send test notification.");
  }
}
