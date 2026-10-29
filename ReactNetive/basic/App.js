import React, { useState, useEffect, useRef } from "react";
import { View, Text, Button, Image, StyleSheet } from "react-native";
import { Camera } from "expo-camera";

console.log(Camera);

function CameraApp() {
  const [hasPermission, setHasPermission] = useState(null);
  const [photo, setPhoto] = useState(null);
  const cameraRef = useRef(null);

  // Request camera permission
  useEffect(() => {
    (async () => {
      const { status } = await Camera.requestCameraPermissionsAsync();
      setHasPermission(status === "granted");
    })();
  }, []);

  // Take a picture
  const takePicture = async () => {
    if (cameraRef.current) {
      try {
        const photoData = await cameraRef.current.takePictureAsync();
        setPhoto(photoData.uri);
      } catch (error) {
        console.error("Error taking picture: ", error);
      }
    }
  };

  if (hasPermission === null) {
    return <Text>Requesting camera permission...</Text>;
  }
  if (hasPermission === false) {
    return <Text>No access to the camera</Text>;
  }

  return (
    <View style={styles.container}>
      {!photo ? (
        <>
          <Camera ref={cameraRef} style={styles.camera} />
          <View style={styles.buttonContainer}>
            <Button title="Take Photo" onPress={takePicture} />
          </View>
        </>
      ) : (
        <>
          <Image source={{ uri: photo }} style={styles.image} />
          <View style={styles.buttonContainer}>
            <Button title="Retake" onPress={() => setPhoto(null)} />
          </View>
        </>
      )}
    </View>
  );
}

export default function App() {
  return <CameraApp />;
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  camera: {
    flex: 1,
  },
  image: {
    flex: 1,
    resizeMode: "cover",
  },
  buttonContainer: {
    flexDirection: "row",
    justifyContent: "space-around",
    margin: 20,
  },
});
