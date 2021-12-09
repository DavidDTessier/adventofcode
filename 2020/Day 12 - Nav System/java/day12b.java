package org.ea;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Puzzle12b {
    public static void main(String[] args) {
        String debug = "";
        try {
            BufferedReader br = new BufferedReader(new FileReader("../input/input.dat"));

            Pattern p = Pattern.compile("([NEWSLRF])(\\d+)");

            String line;
            while ((line = br.readLine()) != null) {
                debug = line;
                Matcher m = p.matcher(line);
                if (m.find()) {
                    handleInstruction(m.group(1), Integer.parseInt(m.group(2)));
                }
            }

            System.out.println("Manhattan Distance: " + (Math.abs(x) + Math.abs(y)));

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

    private static int wx = 10;
    private static int wy = -1;

    private static void handleInstruction(String inst, int value) {
        switch (inst) {
            case "N":
                wy -= value;
                break;
            case "S":
                wy += value;
                break;
            case "E":
                wx += value;
                break;
            case "W":
                wx -= value;
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
        if (value < 0) {
            value += 360;
        }

        double radians = Math.toRadians(value);

        float newx = (float)(wx * Math.cos(radians) - wy * Math.sin(radians));
        float newy = (float)(wx * Math.sin(radians) + wy * Math.cos(radians));

        wx = (int)newx;
        wy = (int)newy;
    }

    private static void moveForward(int value) {
        y += wy * value;
        x += wx * value;
    }
}