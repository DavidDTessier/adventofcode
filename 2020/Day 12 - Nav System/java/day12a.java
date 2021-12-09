package org.ea;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Puzzle12a {
    public static void main(String[] args) {
        String debug = "";
        try {
            BufferedReader br = new BufferedReader(new FileReader("../input/input.dat"));

            Pattern p = Pattern.compile("([NEWSLRF])(\\d+)");

            String line;
            while ((line = br.readLine()) != null) {
                debug = line;
                System.out.print(debug+"\n");
                Matcher m = p.matcher(line);
                if (m.find()) {
                    handleInstruction(m.group(1), Integer.parseInt(m.group(2)));
                }
            }
            System.out.println(x);
            System.out.println(y);
            System.out.println(direction);
            System.out.println("Manhattan Distance (JAVA): " + (Math.abs(x) + Math.abs(y)));

        } catch (Exception e) {
            System.out.println("Debug: " + debug);
            e.printStackTrace();
        }
    }

    private static final int NORTH = 0;
    private static final int EAST = 90;
    private static final int SOUTH = 180;
    private static final int WEST = 270;
    private static int direction = 90;

    private static int x = 0;
    private static int y = 0;

    private static void handleInstruction(String inst, int value) {
        switch (inst) {
            case "N":
                y -= value;
                break;
            case "S":
                y += value;
                break;
            case "E":
                x += value;
                break;
            case "W":
                x -= value;
                break;
            case "F":
                moveForward(value);
                break;
            case "L":
                changeDirection(-value);
                break;
            case "R":
                changeDirection(value);
                break;
        }
    }

    private static void changeDirection(int value) {
        direction += value;
        if (direction < 0) {
            direction += 360;
        }
        if (direction > 359) {
            direction -= 360;
        }
    }

    private static void moveForward(int value) {
        switch (direction) {
            case NORTH:
                y -= value;
                break;
            case SOUTH:
                y += value;
                break;
            case EAST:
                x += value;
                break;
            case WEST:
                x -= value;
                break;
        }
    }
}